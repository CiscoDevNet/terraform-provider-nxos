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
	"reflect"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/netascode/go-nxos"
)

// testSubtreeDevice returns a device backed by a fake NX-API server in the test process running the given version,
// which rejects queries containing any of the unknown classes. Like NX-OS, only the first unknown class of a query is
// reported. It also returns the number of rejected queries and the queried class lists.
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
				if slices.Contains(unknown, c) {
					failed.Add(1)
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, `{"imdata":[{"error":{"attributes":{"code":"17","text":"Unknown class %s"}}}]}`, c)
					return
				}
			}
			fmt.Fprint(w, `{"imdata":[{"a":{"attributes":{"dn":"sys/a"}}}]}`)
		}
	}))
	t.Cleanup(server.Close)
	client, _ := nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))
	return &NxosProviderDataDevice{Client: client, Managed: true}, &failed, &queried
}

func TestNxos_GetSubtree(t *testing.T) {
	cases := []struct {
		name            string
		version         string
		unknown         []string
		classes         []string
		minVersions     map[string]string
		calls           int
		concurrent      bool
		expectedFailed  int32
		expectedQueries []string
	}{
		{
			// Unknown classes are removed from the query and remembered, so the second call fails no request
			name:            "unknown classes",
			version:         "10.4(3)",
			unknown:         []string{"c", "e"},
			classes:         []string{"b", "c", "d", "e"},
			calls:           2,
			expectedFailed:  2,
			expectedQueries: []string{"b,c,d,e", "b,d,e", "b,d"},
		},
		{
			name:            "all classes unknown",
			version:         "10.4(3)",
			unknown:         []string{"b"},
			classes:         []string{"b"},
			calls:           1,
			expectedFailed:  1,
			expectedQueries: []string{"", "b"},
		},
		{
			// Classes requiring a newer version are not queried at all
			name:            "minimum versions",
			version:         "10.5(6)",
			unknown:         []string{"d"},
			classes:         []string{"b", "c", "d"},
			minVersions:     map[string]string{"c": "10.6(1)"},
			calls:           1,
			expectedFailed:  1,
			expectedQueries: []string{"b", "b,d"},
		},
		{
			name:            "concurrent calls",
			version:         "10.4(3)",
			unknown:         []string{"c", "e"},
			classes:         []string{"b", "c", "d", "e"},
			calls:           10,
			concurrent:      true,
			expectedQueries: []string{"b,c,d,e", "b,d,e", "b,d"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			device, failed, queried := testSubtreeDevice(t, c.version, c.unknown...)
			classes := slices.Clone(c.classes)
			get := func() {
				res, err := device.GetSubtree(context.Background(), "sys/a", classes, c.minVersions)
				if err != nil || !res.Exists() {
					t.Errorf("Unexpected error: %v", err)
				}
			}
			var wg sync.WaitGroup
			for i := 0; i < c.calls; i++ {
				if c.concurrent {
					wg.Add(1)
					go func() { defer wg.Done(); get() }()
				} else {
					get()
				}
			}
			wg.Wait()

			if !reflect.DeepEqual(classes, c.classes) {
				t.Fatalf("Error matching classes and input, the input was modified: %v vs %v", classes, c.classes)
			}
			if !c.concurrent && failed.Load() != c.expectedFailed {
				t.Fatalf("Error matching failed requests and expected: %d vs %d", failed.Load(), c.expectedFailed)
			}
			var queries []string
			queried.Range(func(key, _ any) bool { queries = append(queries, key.(string)); return true })
			slices.Sort(queries)
			expected := slices.Clone(c.expectedQueries)
			slices.Sort(expected)
			if c.concurrent {
				// Concurrent calls may not all see the unknown classes yet, but must end up with the filtered query
				if !slices.Contains(queries, expected[0]) {
					t.Fatalf("Error matching queries and expected: %v does not contain %q", queries, expected[0])
				}
			} else if !reflect.DeepEqual(queries, expected) {
				t.Fatalf("Error matching queries and expected: %v vs %v", queries, expected)
			}
		})
	}
}

func TestNxos_GetSubtreeErrors(t *testing.T) {
	cases := []struct {
		name    string
		classes []string
	}{
		{name: "unknown class not part of the query", classes: []string{"x"}},
		{name: "other error code", classes: []string{"z"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			device, _, _ := testSubtreeDevice(t, "10.4(3)")
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

			if _, err := device.GetSubtree(context.Background(), "sys/a", c.classes, nil); err == nil {
				t.Fatalf("Expected an error")
			}
			for _, class := range []string{"x", "y", "z"} {
				if device.isUnsupportedClass(class) {
					t.Fatalf("Error matching unsupported classes and expected: %q is marked as unsupported", class)
				}
			}
		})
	}
}
