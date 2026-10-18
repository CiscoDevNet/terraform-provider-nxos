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
	"reflect"
	"strings"
	"testing"

	"github.com/netascode/go-nxos"
)

// testVersionClient returns a client for a fake NX-API server in the test process reporting the given version.
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

// testNoRequestsClient returns a client for a fake NX-API server in the test process that fails the test on any request.
func testNoRequestsClient(t *testing.T) *nxos.Client {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request: %s", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)
	client, _ := nxos.NewClient(server.URL, "usr", "pwd", true, nxos.MaxRetries(0))
	return client
}

func TestNxos_SubtreeClasses(t *testing.T) {
	classes := []string{"a", "b", "c"}

	cases := []struct {
		name        string
		client      func(t *testing.T) *nxos.Client
		minVersions map[string]string
		expected    []string
	}{
		{name: "older device", client: func(t *testing.T) *nxos.Client { return testVersionClient(t, "10.5(6)") }, minVersions: map[string]string{"b": "10.6(1)", "c": "10.5(1)"}, expected: []string{"a", "c"}},
		{name: "newer device", client: func(t *testing.T) *nxos.Client { return testVersionClient(t, "10.6(1)") }, minVersions: map[string]string{"b": "10.6(1)", "c": "10.5(1)"}, expected: []string{"a", "b", "c"}},
		{name: "unknown version", client: func(t *testing.T) *nxos.Client { return testVersionClient(t, "") }, minVersions: map[string]string{"b": "10.6(1)", "c": "10.5(1)"}, expected: []string{"a", "b", "c"}},
		{name: "no minimum versions", client: testNoRequestsClient, minVersions: map[string]string{}, expected: []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := SubtreeClasses(context.Background(), c.client(t), classes, c.minVersions); !reflect.DeepEqual(out, c.expected) {
				t.Fatalf("Error matching output and expected: %v vs %v", out, c.expected)
			}
		})
	}
}

func TestNxos_CheckMinVersion(t *testing.T) {
	checks := map[string]MinVersion{
		"b":      {Version: "10.6(1)", Path: "list_b"},
		"a.attr": {Version: "10.6(1)", Path: "attr"},
	}
	older := func(t *testing.T) *nxos.Client { return testVersionClient(t, "10.5(6)") }

	cases := []struct {
		name   string
		client func(t *testing.T) *nxos.Client
		checks map[string]MinVersion
		body   string
		errors int
		detail string
	}{
		{name: "unsupported child", client: older, checks: checks, body: `{"a":{"attributes":{},"children":[{"b":{"attributes":{"id":"1"}}}]}}`, errors: 1},
		{name: "unsupported attribute", client: older, checks: checks, body: `{"a":{"attributes":{"attr":"x"}}}`, errors: 1, detail: "Attribute `attr` requires NX-OS 10.6(1) or later, but device 'leaf1' runs NX-OS 10.5(6)."},
		{name: "both", client: older, checks: checks, body: `{"a":{"attributes":{"attr":"x"},"children":[{"b":{"attributes":{}}}]}}`, errors: 2},
		{name: "deleted child", client: older, checks: checks, body: `{"a":{"attributes":{},"children":[{"b":{"attributes":{"status":"deleted"}}}]}}`, errors: 0},
		{name: "unset attribute", client: older, checks: checks, body: `{"a":{"attributes":{"attr":"DME_UNSET_PROPERTY_MARKER"}}}`, errors: 0},
		{name: "supported only", client: older, checks: checks, body: `{"a":{"attributes":{"other":"x"},"children":[{"c":{"attributes":{}}}]}}`, errors: 0},
		{name: "newer device", client: func(t *testing.T) *nxos.Client { return testVersionClient(t, "10.6(4)") }, checks: checks, body: `{"a":{"attributes":{"attr":"x"}}}`, errors: 0},
		{name: "unknown version", client: func(t *testing.T) *nxos.Client { return testVersionClient(t, "") }, checks: checks, body: `{"a":{"attributes":{"attr":"x"}}}`, errors: 0},
		{name: "no minimum versions", client: testNoRequestsClient, checks: map[string]MinVersion{}, body: `{"a":{"attributes":{"x":"y"}}}`, errors: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			diags := CheckMinVersion(context.Background(), c.client(t), "leaf1", c.body, c.checks)
			if diags.ErrorsCount() != c.errors {
				t.Fatalf("Error matching error count and expected: %d vs %d: %v", diags.ErrorsCount(), c.errors, diags)
			}
			if c.detail != "" && diags[0].Detail() != c.detail {
				t.Fatalf("Error matching error detail and expected: %q vs %q", diags[0].Detail(), c.detail)
			}
		})
	}
}
