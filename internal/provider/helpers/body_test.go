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

func TestNxos_FindChildPath(t *testing.T) {
	body := `{"a":{"children":[{"x":{"attributes":{}}},{"b":{"attributes":{}}}]}}`

	cases := []struct {
		name       string
		parentPath string
		className  string
		expected   string
		resolves   bool
	}{
		{name: "existing child", parentPath: "a.children", className: "b", expected: "a.children.1.b", resolves: true},
		{name: "missing child", parentPath: "a.children", className: "c", expected: "a.children.2.c", resolves: false},
		{name: "missing parent", parentPath: "n.children", className: "c", expected: "n.children.0.c", resolves: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := FindChildPath(body, c.parentPath, c.className)
			if out != c.expected {
				t.Fatalf("Error matching output and expected: %q vs %q", out, c.expected)
			}
			if gjson.Get(body, out).Exists() != c.resolves {
				t.Fatalf("Error matching resolution of %q and expected: %v", out, c.resolves)
			}
		})
	}
}

func TestNxos_EnsureChildPath(t *testing.T) {
	cases := []struct {
		name         string
		body         string
		ensure       func(body *string) string
		expectedPath string
		expectedBody string
	}{
		{
			// A missing child is appended after its siblings, which stay separate objects
			name:         "missing child",
			body:         `{"a":{"children":[{"x":{"attributes":{}}}]}}`,
			ensure:       func(body *string) string { return EnsureChildPath(body, "a.children", "b") },
			expectedPath: "a.children.1.b",
			expectedBody: `{"a":{"children":[{"x":{"attributes":{}}},{"b":{"attributes":{}}}]}}`,
		},
		{
			// An existing child is reused, not duplicated
			name:         "existing child",
			body:         `{"a":{"children":[{"x":{"attributes":{}}},{"b":{"attributes":{}}}]}}`,
			ensure:       func(body *string) string { return EnsureChildPath(body, "a.children", "b") },
			expectedPath: "a.children.1.b",
			expectedBody: `{"a":{"children":[{"x":{"attributes":{}}},{"b":{"attributes":{}}}]}}`,
		},
		{
			// A nested ensure creates the whole chain below a missing parent
			name: "missing parent",
			body: `{"r":{"children":[]}}`,
			ensure: func(body *string) string {
				return EnsureChildPath(body, EnsureChildPath(body, "r.children", "p")+".children", "q")
			},
			expectedPath: "r.children.0.p.children.0.q",
			expectedBody: `{"r":{"children":[{"p":{"attributes":{},"children":[{"q":{"attributes":{}}}]}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			body := c.body
			if out := c.ensure(&body); out != c.expectedPath {
				t.Fatalf("Error matching output and expected: %q vs %q", out, c.expectedPath)
			}
			if body != c.expectedBody {
				t.Fatalf("Error matching body and expected:\n%s\nvs\n%s", body, c.expectedBody)
			}
		})
	}
}

func TestNxos_SplitSeparateRequestBody(t *testing.T) {
	udf := SeparateRequest{Attributes: map[string]map[string]string{
		"platformUdf": {"qualifyRacl": "DME_UNSET_PROPERTY_MARKER", "qualifyVacl": "DME_UNSET_PROPERTY_MARKER"},
	}}
	vpc := SeparateRequest{
		Attributes: map[string]map[string]string{
			"vpcDom": {"l3PeerRouter": "DME_UNSET_PROPERTY_MARKER", "peerGWExcludeVLAN": "DME_UNSET_PROPERTY_MARKER"},
		},
		Keys: map[string][]string{"vpcDom": {"id"}},
	}

	cases := []struct {
		name             string
		body             string
		separate         SeparateRequest
		expectedMain     string
		expectedSeparate string
		expectedReset    string
	}{
		{
			// Objects of separate classes are sent last, subtrees of deleted objects stay in the main body
			name: "separate classes",
			body: `{"interfaceEntity":{"attributes":{},"children":[` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"l1VlanMapping":{"attributes":{"Enabled":"yes"}}},{"nwRtVrfMbr":{"attributes":{"tDn":"sys/inst-default"}}}]}},` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/2","mode":"access"}}},` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/3","status":"deleted"},"children":[{"l1VlanMapping":{"attributes":{}}}]}}]}}`,
			separate: SeparateRequest{Classes: []string{"l1VlanMapping"}},
			expectedMain: `{"interfaceEntity":{"attributes":{},"children":[` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"nwRtVrfMbr":{"attributes":{"tDn":"sys/inst-default"}}}]}},` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/2","mode":"access"}}},` +
				`{"l1PhysIf":{"attributes":{"id":"eth1/3","status":"deleted"},"children":[{"l1VlanMapping":{"attributes":{}}}]}}]}}`,
			expectedSeparate: `{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1","mode":"trunk"},"children":[{"l1VlanMapping":{"attributes":{"Enabled":"yes"}}}]}}]}}`,
		},
		{
			name:         "without separate classes",
			body:         `{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1"}}}]}}`,
			separate:     SeparateRequest{Classes: []string{"l1VlanMapping"}},
			expectedMain: `{"interfaceEntity":{"attributes":{},"children":[{"l1PhysIf":{"attributes":{"id":"eth1/1"}}}]}}`,
		},
		{
			// Separate attributes are sent last, those of deleted objects are reset first with only the keys of the parents
			name: "separate attributes",
			body: `{"platformEntity":{"attributes":{"descr":"x"},"children":[` +
				`{"platformTcamTemplate":{"attributes":{"name":"TPL1"}}},` +
				`{"platformUdfEntity":{"attributes":{},"children":[` +
				`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start","qualifyRacl":"true"}}},` +
				`{"platformUdf":{"attributes":{"name":"UDF2","offsetBase":"packet-start"}}},` +
				`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"deleted"}}}]}}]}}`,
			separate: udf,
			expectedMain: `{"platformEntity":{"attributes":{"descr":"x"},"children":[` +
				`{"platformTcamTemplate":{"attributes":{"name":"TPL1"}}},` +
				`{"platformUdfEntity":{"attributes":{},"children":[` +
				`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start"}}},` +
				`{"platformUdf":{"attributes":{"name":"UDF2","offsetBase":"packet-start"}}},` +
				`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"deleted"}}}]}}]}}`,
			expectedSeparate: `{"platformEntity":{"attributes":{"descr":"x"},"children":[{"platformUdfEntity":{"attributes":{},"children":[` +
				`{"platformUdf":{"attributes":{"name":"UDF1","offsetBase":"packet-start","qualifyRacl":"true"}}}]}}]}}`,
			expectedReset: `{"platformEntity":{"attributes":{},"children":[{"platformUdfEntity":{"attributes":{},"children":[` +
				`{"platformUdf":{"attributes":{"rn":"udf-UDF3","status":"modified","qualifyRacl":"DME_UNSET_PROPERTY_MARKER","qualifyVacl":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}`,
		},
		{
			name:             "separate attributes of the root object",
			body:             `{"platformUdf":{"attributes":{"name":"UDF1","qualifyVacl":"false"}}}`,
			separate:         udf,
			expectedMain:     `{"platformUdf":{"attributes":{"name":"UDF1"}}}`,
			expectedSeparate: `{"platformUdf":{"attributes":{"name":"UDF1","qualifyVacl":"false"}}}`,
		},
		{
			name:         "without separate attributes",
			body:         `{"platformEntity":{"attributes":{},"children":[{"platformUdfEntity":{"attributes":{},"children":[{"platformUdf":{"attributes":{"name":"UDF1"}}}]}}]}}`,
			separate:     udf,
			expectedMain: `{"platformEntity":{"attributes":{},"children":[{"platformUdfEntity":{"attributes":{},"children":[{"platformUdf":{"attributes":{"name":"UDF1"}}}]}}]}}`,
		},
		{
			// A removed separate attribute is reset first with only the keys of its object, the others are sent last
			name: "removed separate attributes",
			body: `{"vpcEntity":{"attributes":{},"children":[{"vpcInst":{"attributes":{},"children":[` +
				`{"vpcDom":{"attributes":{"id":"100","peerGw":"DME_UNSET_PROPERTY_MARKER","l3PeerRouter":"DME_UNSET_PROPERTY_MARKER","peerGWExcludeVLAN":"1-2"}}}]}}]}}`,
			separate: vpc,
			expectedMain: `{"vpcEntity":{"attributes":{},"children":[{"vpcInst":{"attributes":{},"children":[` +
				`{"vpcDom":{"attributes":{"id":"100","peerGw":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}`,
			expectedSeparate: `{"vpcEntity":{"attributes":{},"children":[{"vpcInst":{"attributes":{},"children":[` +
				`{"vpcDom":{"attributes":{"id":"100","peerGw":"DME_UNSET_PROPERTY_MARKER","peerGWExcludeVLAN":"1-2"}}}]}}]}}`,
			expectedReset: `{"vpcEntity":{"attributes":{},"children":[{"vpcInst":{"attributes":{},"children":[` +
				`{"vpcDom":{"attributes":{"id":"100","l3PeerRouter":"DME_UNSET_PROPERTY_MARKER","status":"modified"}}}]}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mainBody, separateBody, resetBody := SplitSeparateRequestBody(c.body, c.separate)
			if mainBody != c.expectedMain {
				t.Fatalf("Error matching main body and expected:\n%s\nvs\n%s", mainBody, c.expectedMain)
			}
			if separateBody != c.expectedSeparate {
				t.Fatalf("Error matching separate body and expected:\n%s\nvs\n%s", separateBody, c.expectedSeparate)
			}
			if resetBody != c.expectedReset {
				t.Fatalf("Error matching reset body and expected:\n%s\nvs\n%s", resetBody, c.expectedReset)
			}
		})
	}
}
