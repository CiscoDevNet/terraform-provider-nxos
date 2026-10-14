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

func TestSplitSeparateRequestBody(t *testing.T) {
	body := `{"interfaceEntity":{"attributes":{},"children":[` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"l1VlanMapping":{"attributes":{"Enabled":"yes"}}},{"nwRtVrfMbr":{"attributes":{"tDn":"sys/inst-default"}}}]}},` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/2","mode":"access"}}},` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/3","status":"deleted"},"children":[{"l1VlanMapping":{"attributes":{}}}]}}]}}`
	mainBody, separateBody, resetBody := SplitSeparateRequestBody(body, SeparateRequest{Classes: []string{"l1VlanMapping"}})
	wantMain := `{"interfaceEntity":{"attributes":{},"children":[` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"nwRtVrfMbr":{"attributes":{"tDn":"sys/inst-default"}}}]}},` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/2","mode":"access"}}},` +
		`{"l1PhysIf":{"attributes":{"id":"eth1/3","status":"deleted"},"children":[{"l1VlanMapping":{"attributes":{}}}]}}]}}`
	wantSeparate := `{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"l1VlanMapping":{"attributes":{"Enabled":"yes"}}}]}}]}}`
	if mainBody != wantMain {
		t.Errorf("main body:\n got %s\nwant %s", mainBody, wantMain)
	}
	if separateBody != wantSeparate {
		t.Errorf("separate body:\n got %s\nwant %s", separateBody, wantSeparate)
	}
	if resetBody != "" {
		t.Errorf("reset body must be empty without separate attributes, got %s", resetBody)
	}

	mainBody, separateBody, resetBody = SplitSeparateRequestBody(`{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1"}}}]}}`, SeparateRequest{Classes: []string{"l1VlanMapping"}})
	if separateBody != "" || resetBody != "" || mainBody != `{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1"}}}]}}` {
		t.Errorf("body without separate classes must stay unchanged, got main %s separate %s reset %s", mainBody, separateBody, resetBody)
	}
}

func TestSplitSeparateRequestBodyAttributes(t *testing.T) {
	separate := SeparateRequest{Attributes: map[string]map[string]string{
		"platformUdf": {"qualifyRacl": "DME_UNSET_PROPERTY_MARKER", "qualifyVacl": "DME_UNSET_PROPERTY_MARKER"},
	}}
	body := `{"platformEntity":{"attributes":{"descr":"x"},"children":[` +
		`{"platformTcamTemplate":{"attributes":{"name":"TPL1"}}},` +
		`{"platformUdfEntity":{"attributes":{},"children":[` +
		`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start","qualifyRacl":"true"}}},` +
		`{"platformUdf":{"attributes":{"name":"UDF2","offsetBase":"packet-start"}}},` +
		`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"deleted"}}}]}}]}}`
	mainBody, separateBody, resetBody := SplitSeparateRequestBody(body, separate)
	wantMain := `{"platformEntity":{"attributes":{"descr":"x"},"children":[` +
		`{"platformTcamTemplate":{"attributes":{"name":"TPL1"}}},` +
		`{"platformUdfEntity":{"attributes":{},"children":[` +
		`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start"}}},` +
		`{"platformUdf":{"attributes":{"name":"UDF2","offsetBase":"packet-start"}}},` +
		`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"deleted"}}}]}}]}}`
	wantSeparate := `{"platformEntity":{"attributes":{"descr":"x"},"children":[{"platformUdfEntity":{"attributes":{},"children":[` +
		`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start","qualifyRacl":"true"}}}]}}]}}`
	wantReset := `{"platformEntity":{"attributes":{"descr":"x"},"children":[{"platformUdfEntity":{"attributes":{},"children":[` +
		`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"modified","qualifyRacl":"DME_UNSET_PROPERTY_MARKER","qualifyVacl":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}`
	if mainBody != wantMain {
		t.Errorf("main body:\n got %s\nwant %s", mainBody, wantMain)
	}
	if separateBody != wantSeparate {
		t.Errorf("separate body:\n got %s\nwant %s", separateBody, wantSeparate)
	}
	if resetBody != wantReset {
		t.Errorf("reset body:\n got %s\nwant %s", resetBody, wantReset)
	}

	// Separate attributes of the root object itself
	mainBody, separateBody, resetBody = SplitSeparateRequestBody(`{"platformUdf":{"attributes":{"name":"UDF1","qualifyVacl":"false"}}}`, separate)
	if mainBody != `{"platformUdf":{"attributes":{"name":"UDF1"}}}` || separateBody != `{"platformUdf":{"attributes":{"name":"UDF1","qualifyVacl":"false"}}}` || resetBody != "" {
		t.Errorf("root object: got main %s separate %s reset %s", mainBody, separateBody, resetBody)
	}

	// A body without separate attributes stays unchanged
	unchanged := `{"platformEntity":{"attributes":{},"children":[{"platformUdfEntity":{"attributes":{},"children":[{"platformUdf":{"attributes":{"name":"UDF1"}}}]}}]}}`
	mainBody, separateBody, resetBody = SplitSeparateRequestBody(unchanged, separate)
	if mainBody != unchanged || separateBody != "" || resetBody != "" {
		t.Errorf("body without separate attributes must stay unchanged, got main %s separate %s reset %s", mainBody, separateBody, resetBody)
	}
}
