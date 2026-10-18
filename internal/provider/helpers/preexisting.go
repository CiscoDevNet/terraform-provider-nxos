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
	"encoding/json"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/tidwall/gjson"
)

// preexistingKey is the private state key holding the DNs of the managed objects that already existed on the device
// before the resource managed them. The `reset` deletion policy resets these objects and deletes all others.
const preexistingKey = "preexisting"

// GetPreexisting returns the DNs of the managed objects that already existed on the device before the resource
// managed them. It returns nil if nothing was recorded, e.g. for state written by provider versions before the
// deletion policy was introduced.
func GetPreexisting(ctx context.Context, gk GetKeyer) (map[string]bool, diag.Diagnostics) {
	v, diags := gk.GetKey(ctx, preexistingKey)
	if diags.HasError() || len(v) == 0 {
		return nil, diags
	}
	var dns []string
	if err := json.Unmarshal(v, &dns); err != nil {
		diags.AddError("Private State Error", "Failed to parse the preexisting objects: "+err.Error())
		return nil, diags
	}
	preexisting := make(map[string]bool, len(dns))
	for _, dn := range dns {
		preexisting[dn] = true
	}
	return preexisting, diags
}

// SetPreexisting checks the respDiags and if they are error-free it records the DNs of the managed objects that
// already existed on the device before the resource managed them. It appends its own results to respDiags.
func SetPreexisting(ctx context.Context, sk SetKeyer, preexisting map[string]bool, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}
	dns := make([]string, 0, len(preexisting))
	for dn := range preexisting {
		dns = append(dns, dn)
	}
	slices.Sort(dns)
	v, err := json.Marshal(dns)
	if err != nil {
		respDiags.AddError("Private State Error", "Failed to store the preexisting objects: "+err.Error())
		return
	}
	respDiags.Append(sk.SetKey(ctx, preexistingKey, v)...)
}

// Tree holds the objects of a subtree query result by their DN.
type Tree map[string]TreeObject

// TreeObject is an object of a Tree.
type TreeObject struct {
	Class  string
	Parent string
}

// ReadTree returns the objects in the result of a subtree query of dn.
func ReadTree(res gjson.Result, dn string) Tree {
	tree := Tree{}
	var walk func(obj gjson.Result, dn, parent string)
	walk = func(obj gjson.Result, dn, parent string) {
		obj.ForEach(func(class, body gjson.Result) bool {
			tree[dn] = TreeObject{Class: class.String(), Parent: parent}
			body.Get("children").ForEach(func(_, child gjson.Result) bool {
				child.ForEach(func(_, childBody gjson.Result) bool {
					if rn := childBody.Get("attributes.rn").String(); rn != "" {
						walk(child, dn+"/"+rn, dn)
					}
					return true
				})
				return true
			})
			return true
		})
	}
	if res.Exists() {
		walk(res, dn, "")
	}
	return tree
}

// Dns returns the DNs of all objects of the tree.
func (t Tree) Dns() map[string]bool {
	dns := make(map[string]bool, len(t))
	for dn := range t {
		dns[dn] = true
	}
	return dns
}

// HasUnmanagedEntries returns true if the tree holds a list entry below dn that is not part of managed, e.g. an entry
// configured by another resource. Entries below managed list entries do not count, as they belong to those entries.
func (t Tree) HasUnmanagedEntries(dn string, managed, listClasses map[string]bool) bool {
	for e, obj := range t {
		if managed[e] || !listClasses[obj.Class] {
			continue
		}
		withinManagedEntry := false
		for p := obj.Parent; p != ""; p = t[p].Parent {
			if p == dn {
				if !withinManagedEntry {
					return true
				}
				break
			}
			if managed[p] && listClasses[t[p].Class] {
				withinManagedEntry = true
			}
		}
	}
	return false
}

// FilterExisting returns the DNs of dns that are part of existing.
func FilterExisting(dns, existing map[string]bool) map[string]bool {
	filtered := map[string]bool{}
	for dn := range dns {
		if existing[dn] {
			filtered[dn] = true
		}
	}
	return filtered
}

var (
	cr resource.CreateResponse
	ur resource.UpdateResponse
	dq resource.DeleteRequest
	_  SetKeyer = cr.Private
	_  SetKeyer = ur.Private
	_  GetKeyer = dq.Private
)
