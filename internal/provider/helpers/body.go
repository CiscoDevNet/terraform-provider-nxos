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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// FindChildPath returns the path of the child node of class className in the
// children array at parentPath, e.g. "a.children.1.b". The array is searched
// by class name, so the result does not depend on how many siblings precede
// the child. If there is no such child, the returned path points one past the
// end of the array and therefore resolves to nothing when queried.
func FindChildPath(body, parentPath, className string) string {
	children := gjson.Get(body, parentPath).Array()
	for i, child := range children {
		if child.Get(className).Exists() {
			return parentPath + "." + strconv.Itoa(i) + "." + className
		}
	}
	return parentPath + "." + strconv.Itoa(len(children)) + "." + className
}

// EnsureChildPath is like FindChildPath but appends an empty child node of
// class className to the array at parentPath if none exists yet, so that the
// returned path can be written to.
func EnsureChildPath(body *string, parentPath, className string) string {
	path := FindChildPath(*body, parentPath, className)
	if !gjson.Get(*body, path).Exists() {
		*body, _ = sjson.SetRaw(*body, path+".attributes", "{}")
	}
	return path
}

// SeparateRequest describes the parts of a body that the device only accepts
// once their object already exists.
type SeparateRequest struct {
	// Classes lists classes whose objects are sent in a second request after the
	// rest of the body.
	Classes []string
	// Attributes maps a class to attributes that are sent in a second request,
	// once the object exists, together with the value that resets each of them.
	// When they are removed, or before an object of the class is deleted, these
	// attributes are reset in a request of their own before the rest of the body,
	// because the device rejects removing what they depend on, or deleting the
	// object, while they are set.
	Attributes map[string]map[string]string
	// Keys maps a class to the attributes identifying its objects, which are the
	// only other attributes sent with the reset of separate attributes.
	Keys map[string][]string
}

// SplitSeparateRequestBody splits body into up to three bodies that are sent
// in this order:
//   - a reset body that resets the separate attributes that are removed, and
//     those of objects about to be deleted, marked with status "modified" so
//     that the request fails instead of creating an object that does not exist
//   - the main body without objects of the separate classes and without the
//     separate attributes
//   - a separate body with those objects and attributes
//
// The reset and separate bodies also hold the chain of parent objects leading
// to them, and are empty if there is nothing to send. Subtrees of deleted
// objects stay in the main body.
func SplitSeparateRequestBody(body string, separate SeparateRequest) (string, string, string) {
	classes := make(map[string]bool, len(separate.Classes))
	for _, c := range separate.Classes {
		classes[c] = true
	}
	return splitSeparateRequestNode(gjson.Parse(body), classes, separate.Attributes, separate.Keys)
}

// keyAttributes returns the attributes of attrs that identify the object: its keys and its rn.
func keyAttributes(attrs string, keys []string) string {
	result := "{}"
	for _, name := range append([]string{"rn"}, keys...) {
		if v := gjson.Get(attrs, gjsonEscape(name)); v.Exists() {
			result, _ = sjson.Set(result, gjsonEscape(name), v.String())
		}
	}
	return result
}

func splitSeparateRequestNode(node gjson.Result, classes map[string]bool, attributes map[string]map[string]string, keys map[string][]string) (string, string, string) {
	var className string
	var content gjson.Result
	node.ForEach(func(key, value gjson.Result) bool {
		className, content = key.String(), value
		return false
	})
	if className == "" {
		return node.Raw, "", ""
	}
	attrs := content.Get("attributes").Raw
	if attrs == "" {
		attrs = "{}"
	}
	separateAttrs := attributes[className]
	if content.Get("attributes.status").String() == "deleted" {
		if len(separateAttrs) == 0 {
			return node.Raw, "", ""
		}
		resetAttrs := attrs
		for _, name := range sortedKeys(separateAttrs) {
			resetAttrs, _ = sjson.Set(resetAttrs, gjsonEscape(name), separateAttrs[name])
		}
		resetAttrs, _ = sjson.Set(resetAttrs, "status", "modified")
		return node.Raw, "", `{"` + className + `":{"attributes":` + resetAttrs + `}}`
	}
	mainAttrs := attrs
	separateNodeAttrs := attrs
	resetNodeAttrs := keyAttributes(attrs, keys[className])
	movedAttrs, resetAttrs := false, false
	for _, name := range sortedKeys(separateAttrs) {
		v := gjson.Get(attrs, gjsonEscape(name))
		if !v.Exists() {
			continue
		}
		mainAttrs, _ = sjson.Delete(mainAttrs, gjsonEscape(name))
		if v.String() == separateAttrs[name] {
			// A removed attribute is reset before the rest of the body
			separateNodeAttrs, _ = sjson.Delete(separateNodeAttrs, gjsonEscape(name))
			resetNodeAttrs, _ = sjson.Set(resetNodeAttrs, gjsonEscape(name), v.String())
			resetAttrs = true
		} else {
			movedAttrs = true
		}
	}
	var mainChildren, separateChildren, resetChildren []string
	for _, child := range content.Get("children").Array() {
		var childClassName string
		child.ForEach(func(key, _ gjson.Result) bool {
			childClassName = key.String()
			return false
		})
		if classes[childClassName] {
			separateChildren = append(separateChildren, child.Raw)
			continue
		}
		mainChild, separateChild, resetChild := splitSeparateRequestNode(child, classes, attributes, keys)
		mainChildren = append(mainChildren, mainChild)
		if separateChild != "" {
			separateChildren = append(separateChildren, separateChild)
		}
		if resetChild != "" {
			resetChildren = append(resetChildren, resetChild)
		}
	}
	if !movedAttrs && !resetAttrs && len(separateChildren) == 0 && len(resetChildren) == 0 {
		return node.Raw, "", ""
	}
	mainContent := content.Raw
	if movedAttrs {
		mainContent, _ = sjson.SetRaw(mainContent, "attributes", mainAttrs)
	}
	if len(mainChildren) > 0 {
		mainContent, _ = sjson.SetRaw(mainContent, "children", "["+strings.Join(mainChildren, ",")+"]")
	} else {
		mainContent, _ = sjson.Delete(mainContent, "children")
	}
	mainBody := `{"` + className + `":` + mainContent + `}`
	separateBody := ""
	if movedAttrs || len(separateChildren) > 0 {
		separateContent := `{"attributes":` + separateNodeAttrs
		if len(separateChildren) > 0 {
			separateContent += `,"children":[` + strings.Join(separateChildren, ",") + `]`
		}
		separateBody = `{"` + className + `":` + separateContent + `}}`
	}
	resetBody := ""
	if resetAttrs || len(resetChildren) > 0 {
		resetContent := `{"attributes":` + keyAttributes(attrs, keys[className])
		if resetAttrs {
			resetNodeAttrs, _ = sjson.Set(resetNodeAttrs, "status", "modified")
			resetContent = `{"attributes":` + resetNodeAttrs
		}
		if len(resetChildren) > 0 {
			resetContent += `,"children":[` + strings.Join(resetChildren, ",") + `]`
		}
		resetBody = `{"` + className + `":` + resetContent + `}}`
	}
	return mainBody, separateBody, resetBody
}

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// gjsonEscape escapes the characters that have a special meaning in gjson and
// sjson paths, so that name addresses a single key.
func gjsonEscape(name string) string {
	var b strings.Builder
	for _, c := range name {
		if strings.ContainsRune(`\.*?|#@!:`, c) {
			b.WriteRune('\\')
		}
		b.WriteRune(c)
	}
	return b.String()
}

// SeparateRequestDelay is the time PostSeparate waits between the two requests.
// The device applies some parent changes asynchronously, e.g. an Ethernet
// switchport mode change takes about half a second to become operational.
var SeparateRequestDelay = 2 * time.Second

// PostSeparate posts body to dn and sends the parts described by separate in
// additional requests (see SplitSeparateRequestBody). Some objects are only
// applied by the device if their parent already exists with its final
// configuration, for example VLAN mappings that depend on the switchport mode
// set in the same request. Sent in the same request, or right after the parent
// change, the device accepts them but does not apply them. Some attributes are
// rejected when they are set on an object that is created in the same request,
// and block deleting the object while they are set, for example UDF qualifiers.
//
// An error of the reset request is ignored: it fails if an object to be deleted
// no longer exists, and the main request reports any other problem.
func PostSeparate(client *nxos.Client, dn, body string, separate SeparateRequest) (nxos.Res, error) {
	mainBody, separateBody, resetBody := SplitSeparateRequestBody(body, separate)
	if resetBody != "" {
		_, _ = client.Post(dn, resetBody)
	}
	res, err := client.Post(dn, mainBody)
	if err != nil || separateBody == "" {
		return res, err
	}
	time.Sleep(SeparateRequestDelay)
	return client.Post(dn, separateBody)
}
