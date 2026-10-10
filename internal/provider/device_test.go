// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/netascode/go-nxos"
)

// testSubtreeDevice returns a device backed by a fake NX-API server running the given version, which rejects
// queries containing any of the unknown classes. Like NX-OS, only the first unknown class of a query is reported.
func testSubtreeDevice(t *testing.T, version string, unknown ...string) (*NxosProviderDataDevice, *atomic.Int32, *sync.Map) {
	var failed atomic.Int32
	var queried sync.Map
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/aaaLogin"):
			fmt.Fprint(w, `{"imdata":[{"aaaLogin":{"attributes":{"token":"abc"}}}]}`)
		case strings.HasPrefix(r.URL.Path, "/api/mo/sys/showversion"):
			fmt.Fprintf(w, `{"imdata":[{"sysmgrShowVersion":{"attributes":{"nxosVersion":"%s"}}}]}`, version)
		default:
			classes := r.URL.Query().Get("rsp-subtree-class")
			queried.Store(classes, true)
			for _, c := range strings.Split(classes, ",") {
				for _, u := range unknown {
					if c == u {
						failed.Add(1)
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, `{"imdata":[{"error":{"attributes":{"code":"17","text":"Unknown class %s"}}}]}`, c)
						return
					}
				}
			}
			fmt.Fprint(w, `{"imdata":[{"a":{"attributes":{"dn":"sys/a"}}}]}`)
		}
	}))
	t.Cleanup(server.Close)
	client, _ := nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))
	return &NxosProviderDataDevice{Client: client, Managed: true}, &failed, &queried
}

func TestGetSubtreeUnknownClasses(t *testing.T) {
	ctx := context.Background()
	device, failed, queried := testSubtreeDevice(t, "10.4(3)", "c", "e")
	classes := []string{"b", "c", "d", "e"}

	res, err := device.GetSubtree(ctx, "sys/a", classes, nil)
	if err != nil || !res.Exists() {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := failed.Load(); got != 2 {
		t.Errorf("got %d failed requests, want 2", got)
	}
	if _, ok := queried.Load("b,d"); !ok {
		t.Errorf("expected a query with the unknown classes removed")
	}
	if strings.Join(classes, ",") != "b,c,d,e" {
		t.Errorf("input classes modified: %v", classes)
	}

	// Unknown classes are remembered, no further failed requests
	if _, err := device.GetSubtree(ctx, "sys/a", classes, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := failed.Load(); got != 2 {
		t.Errorf("got %d failed requests after caching, want 2", got)
	}
}

func TestGetSubtreeAllClassesUnknown(t *testing.T) {
	device, _, queried := testSubtreeDevice(t, "10.4(3)", "b")
	if _, err := device.GetSubtree(context.Background(), "sys/a", []string{"b"}, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := queried.Load(""); !ok {
		t.Errorf("expected a query without class filter")
	}
}

func TestGetSubtreeOtherErrors(t *testing.T) {
	// Unknown class that was not part of the query is not ignored
	device, _, _ := testSubtreeDevice(t, "10.4(3)", "b")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/aaaLogin") {
			fmt.Fprint(w, `{"imdata":[{"aaaLogin":{"attributes":{"token":"abc"}}}]}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		if r.URL.Query().Get("rsp-subtree-class") == "x" {
			fmt.Fprint(w, `{"imdata":[{"error":{"attributes":{"code":"17","text":"Unknown class y"}}}]}`)
		} else {
			fmt.Fprint(w, `{"imdata":[{"error":{"attributes":{"code":"1","text":"Unknown class x"}}}]}`)
		}
	}))
	t.Cleanup(server.Close)
	device.Client, _ = nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))

	if _, err := device.GetSubtree(context.Background(), "sys/a", []string{"x"}, nil); err == nil {
		t.Errorf("unknown class not in query: expected error")
	}
	if _, err := device.GetSubtree(context.Background(), "sys/a", []string{"z"}, nil); err == nil {
		t.Errorf("other error code: expected error")
	}
	if device.isUnsupportedClass("x") || device.isUnsupportedClass("y") || device.isUnsupportedClass("z") {
		t.Errorf("classes wrongly marked as unsupported")
	}
}

func TestGetSubtreeMinVersions(t *testing.T) {
	device, failed, queried := testSubtreeDevice(t, "10.5(6)", "d")
	if _, err := device.GetSubtree(context.Background(), "sys/a", []string{"b", "c", "d"}, map[string]string{"c": "10.6(1)"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := queried.Load("b,d"); !ok {
		t.Errorf("expected first query with min version filtered classes")
	}
	if _, ok := queried.Load("b"); !ok {
		t.Errorf("expected retry with unknown class removed")
	}
	if got := failed.Load(); got != 1 {
		t.Errorf("got %d failed requests, want 1", got)
	}
}

func TestGetSubtreeConcurrent(t *testing.T) {
	device, _, _ := testSubtreeDevice(t, "10.4(3)", "c", "e")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := device.GetSubtree(context.Background(), "sys/a", []string{"b", "c", "d", "e"}, nil); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()
}
