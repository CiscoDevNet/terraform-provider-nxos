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
	"testing"

	"github.com/tidwall/gjson"
)

func TestFindChildPath(t *testing.T) {
	body := `{"a":{"children":[{"x":{"attributes":{}}},{"b":{"attributes":{}}}]}}`

	if got, want := FindChildPath(body, "a.children", "b"), "a.children.1.b"; got != want {
		t.Errorf("existing child: got %q, want %q", got, want)
	}
	missing := FindChildPath(body, "a.children", "c")
	if gjson.Get(body, missing).Exists() {
		t.Errorf("missing child path %q must not resolve", missing)
	}
	if gjson.Get(body, FindChildPath(body, "n.children", "c")).Exists() {
		t.Errorf("child of a missing parent must not resolve")
	}
}

func TestEnsureChildPath(t *testing.T) {
	body := `{"a":{"children":[{"x":{"attributes":{}}}]}}`

	// Missing child is appended after its siblings.
	path := EnsureChildPath(&body, "a.children", "b")
	if want := "a.children.1.b"; path != want {
		t.Fatalf("got path %q, want %q", path, want)
	}
	if !gjson.Get(body, "a.children.0.x").Exists() || !gjson.Get(body, "a.children.1.b.attributes").Exists() {
		t.Fatalf("unexpected body: %s", body)
	}
	if n := len(gjson.Get(body, "a.children").Array()); n != 2 {
		t.Fatalf("expected 2 children, got %d: %s", n, body)
	}
	// Siblings must stay separate objects.
	if gjson.Get(body, "a.children.0.b").Exists() {
		t.Fatalf("child merged into sibling: %s", body)
	}

	// Existing child is reused, not duplicated.
	if again := EnsureChildPath(&body, "a.children", "b"); again != path {
		t.Fatalf("got path %q, want %q", again, path)
	}
	if n := len(gjson.Get(body, "a.children").Array()); n != 2 {
		t.Fatalf("expected 2 children, got %d: %s", n, body)
	}

	// Nested ensure creates the whole chain below a missing parent.
	body = `{"r":{"children":[]}}`
	leaf := EnsureChildPath(&body, EnsureChildPath(&body, "r.children", "p")+".children", "q")
	if want := "r.children.0.p.children.0.q"; leaf != want {
		t.Fatalf("got path %q, want %q", leaf, want)
	}
	if !gjson.Get(body, leaf+".attributes").Exists() {
		t.Fatalf("unexpected body: %s", body)
	}
}
