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
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tidwall/gjson"
)

func TestNxos_ReadTree(t *testing.T) {
	cases := []struct {
		name     string
		res      string
		expected Tree
	}{
		{
			name: "subtree",
			res: `{"slaEntity":{"attributes":{"dn":"sys/sla"},"children":[{"slaInst":{"attributes":{"rn":"inst"},"children":[` +
				`{"slaSender":{"attributes":{"rn":"sender"},"children":[{"slaProbe":{"attributes":{"rn":"operation-2"}}}]}}]}}]}}`,
			expected: Tree{
				"sys/sla":                         {Class: "slaEntity"},
				"sys/sla/inst":                    {Class: "slaInst", Parent: "sys/sla"},
				"sys/sla/inst/sender":             {Class: "slaSender", Parent: "sys/sla/inst"},
				"sys/sla/inst/sender/operation-2": {Class: "slaProbe", Parent: "sys/sla/inst/sender"},
			},
		},
		{
			name:     "missing object",
			res:      ``,
			expected: Tree{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := ReadTree(gjson.Parse(c.res), "sys/sla"); !reflect.DeepEqual(out, c.expected) {
				t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
			}
		})
	}
}

func TestNxos_TreeHasUnmanagedEntries(t *testing.T) {
	tree := Tree{
		"sys/sla":                                  {Class: "slaEntity"},
		"sys/sla/inst":                             {Class: "slaInst", Parent: "sys/sla"},
		"sys/sla/inst/sender":                      {Class: "slaSender", Parent: "sys/sla/inst"},
		"sys/sla/inst/sender/operation-2":          {Class: "slaProbe", Parent: "sys/sla/inst/sender"},
		"sys/sla/inst/sender/operation-2/icmpecho": {Class: "slaIcmpEcho", Parent: "sys/sla/inst/sender/operation-2"},
		"sys/sla/inst/responder":                   {Class: "slaResponder", Parent: "sys/sla/inst"},
	}
	listClasses := map[string]bool{"slaProbe": true}
	sender := "sys/sla/inst/sender"

	cases := []struct {
		name     string
		dn       string
		managed  map[string]bool
		expected bool
	}{
		{name: "entry configured elsewhere", dn: sender, managed: map[string]bool{sender: true}, expected: true},
		{name: "managed entry", dn: sender, managed: map[string]bool{sender: true, sender + "/operation-2": true}, expected: false},
		{name: "no entries below", dn: "sys/sla/inst/responder", managed: map[string]bool{}, expected: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := tree.HasUnmanagedEntries(c.dn, c.managed, listClasses); out != c.expected {
				t.Fatalf("Error matching output and expected: %v vs %v", out, c.expected)
			}
		})
	}
}

func TestNxos_FilterExisting(t *testing.T) {
	cases := []struct {
		name     string
		dns      map[string]bool
		existing map[string]bool
		expected map[string]bool
	}{
		{name: "some exist", dns: map[string]bool{"a": true, "b": true}, existing: map[string]bool{"b": true, "c": true}, expected: map[string]bool{"b": true}},
		{name: "nothing exists", dns: map[string]bool{"a": true}, existing: nil, expected: map[string]bool{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := FilterExisting(c.dns, c.existing); !reflect.DeepEqual(out, c.expected) {
				t.Fatalf("Error matching output and expected: %v vs %v", out, c.expected)
			}
		})
	}
}

// testPrivateState is an in-memory private state.
type testPrivateState map[string][]byte

func (p testPrivateState) GetKey(_ context.Context, key string) ([]byte, diag.Diagnostics) {
	return p[key], nil
}

func (p testPrivateState) SetKey(_ context.Context, key string, value []byte) diag.Diagnostics {
	p[key] = value
	return nil
}

func TestNxos_Preexisting(t *testing.T) {
	cases := []struct {
		name     string
		stored   map[string]bool
		expected map[string]bool
	}{
		{name: "recorded", stored: map[string]bool{"sys/bgp": true, "sys/bgp/inst": true}, expected: map[string]bool{"sys/bgp": true, "sys/bgp/inst": true}},
		{name: "recorded empty", stored: map[string]bool{}, expected: map[string]bool{}},
		{name: "not recorded", stored: nil, expected: nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			private := testPrivateState{}
			if c.stored != nil {
				var diags diag.Diagnostics
				SetPreexisting(context.Background(), private, c.stored, &diags)
				if diags.HasError() {
					t.Fatalf("Unexpected error: %v", diags)
				}
			}
			out, diags := GetPreexisting(context.Background(), private)
			if diags.HasError() {
				t.Fatalf("Unexpected error: %v", diags)
			}
			if !reflect.DeepEqual(out, c.expected) {
				t.Fatalf("Error matching output and expected: %v vs %v", out, c.expected)
			}
		})
	}
}
