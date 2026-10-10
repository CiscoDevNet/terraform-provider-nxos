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

package helpers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/netascode/go-nxos"
)

// testVersionClient returns a client for a fake NX-API server reporting the given version.
// An empty version makes the version request fail.
func testVersionClient(t *testing.T, version string) *nxos.Client {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/aaaLogin"):
			fmt.Fprint(w, `{"imdata":[{"aaaLogin":{"attributes":{"token":"abc"}}}]}`)
		case strings.HasPrefix(r.URL.Path, "/api/mo/sys/showversion") && version != "":
			fmt.Fprintf(w, `{"imdata":[{"sysmgrShowVersion":{"attributes":{"nxosVersion":"%s"}}}]}`, version)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
	t.Cleanup(server.Close)
	client, _ := nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))
	return client
}

func TestSubtreeClasses(t *testing.T) {
	classes := []string{"a", "b", "c"}
	minVersions := map[string]string{"b": "10.6(1)", "c": "10.5(1)"}

	if got, want := strings.Join(SubtreeClasses(context.Background(), testVersionClient(t, "10.5(6)"), classes, minVersions), ","), "a,c"; got != want {
		t.Errorf("older device: got %q, want %q", got, want)
	}
	if got, want := strings.Join(SubtreeClasses(context.Background(), testVersionClient(t, "10.6(1)"), classes, minVersions), ","), "a,b,c"; got != want {
		t.Errorf("newer device: got %q, want %q", got, want)
	}
	if got, want := strings.Join(SubtreeClasses(context.Background(), testVersionClient(t, ""), classes, minVersions), ","), "a,b,c"; got != want {
		t.Errorf("unknown version: got %q, want %q", got, want)
	}
}

func TestCheckMinVersion(t *testing.T) {
	checks := map[string]MinVersion{
		"b":      {Version: "10.6(1)", Path: "list_b"},
		"a.attr": {Version: "10.6(1)", Path: "attr"},
	}
	ctx := context.Background()
	client := testVersionClient(t, "10.5(6)")

	cases := []struct {
		name   string
		body   string
		errors int
	}{
		{"unsupported child", `{"a":{"attributes":{},"children":[{"b":{"attributes":{"id":"1"}}}]}}`, 1},
		{"unsupported attribute", `{"a":{"attributes":{"attr":"x"}}}`, 1},
		{"both", `{"a":{"attributes":{"attr":"x"},"children":[{"b":{"attributes":{}}}]}}`, 2},
		{"deleted child", `{"a":{"attributes":{},"children":[{"b":{"attributes":{"status":"deleted"}}}]}}`, 0},
		{"unset attribute", `{"a":{"attributes":{"attr":"DME_UNSET_PROPERTY_MARKER"}}}`, 0},
		{"supported only", `{"a":{"attributes":{"other":"x"},"children":[{"c":{"attributes":{}}}]}}`, 0},
	}
	for _, c := range cases {
		diags := CheckMinVersion(ctx, client, "leaf1", c.body, checks)
		if diags.ErrorsCount() != c.errors {
			t.Errorf("%s: got %d errors, want %d: %v", c.name, diags.ErrorsCount(), c.errors, diags)
		}
	}

	diags := CheckMinVersion(ctx, client, "leaf1", `{"a":{"attributes":{"attr":"x"}}}`, checks)
	if got := diags[0].Detail(); !strings.Contains(got, "`attr` requires NX-OS 10.6(1)") || !strings.Contains(got, "device 'leaf1' runs NX-OS 10.5(6)") {
		t.Errorf("unexpected error detail: %s", got)
	}

	if diags := CheckMinVersion(ctx, testVersionClient(t, "10.6(4)"), "", `{"a":{"attributes":{"attr":"x"}}}`, checks); diags.HasError() {
		t.Errorf("newer device: unexpected errors: %v", diags)
	}
	if diags := CheckMinVersion(ctx, testVersionClient(t, ""), "", `{"a":{"attributes":{"attr":"x"}}}`, checks); diags.HasError() {
		t.Errorf("unknown version: unexpected errors: %v", diags)
	}
}

func TestMinVersionNoRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)
	client, _ := nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))

	if got, want := strings.Join(SubtreeClasses(context.Background(), client, []string{"a", "b"}, map[string]string{}), ","), "a,b"; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	if diags := CheckMinVersion(context.Background(), client, "", `{"a":{"attributes":{"x":"y"}}}`, map[string]MinVersion{}); diags.HasError() {
		t.Errorf("unexpected errors: %v", diags)
	}
}
