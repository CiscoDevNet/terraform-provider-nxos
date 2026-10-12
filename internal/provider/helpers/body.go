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

// SplitSeparateRequestBody splits body into a main body without any objects of
// the given classes and a second body that holds only those objects, together
// with the chain of parent objects leading to them. The second body is empty if
// body contains no such objects. Subtrees of deleted objects stay in the main
// body.
func SplitSeparateRequestBody(body string, classNames []string) (string, string) {
	classes := make(map[string]bool, len(classNames))
	for _, c := range classNames {
		classes[c] = true
	}
	return splitSeparateRequestNode(gjson.Parse(body), classes)
}

func splitSeparateRequestNode(node gjson.Result, classes map[string]bool) (string, string) {
	var className string
	var content gjson.Result
	node.ForEach(func(key, value gjson.Result) bool {
		className, content = key.String(), value
		return false
	})
	if className == "" || content.Get("attributes.status").String() == "deleted" || !content.Get("children").Exists() {
		return node.Raw, ""
	}
	var mainChildren, separateChildren []string
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
		mainChild, separateChild := splitSeparateRequestNode(child, classes)
		mainChildren = append(mainChildren, mainChild)
		if separateChild != "" {
			separateChildren = append(separateChildren, separateChild)
		}
	}
	if len(separateChildren) == 0 {
		return node.Raw, ""
	}
	mainContent := content.Raw
	if len(mainChildren) > 0 {
		mainContent, _ = sjson.SetRaw(mainContent, "children", "["+strings.Join(mainChildren, ",")+"]")
	} else {
		mainContent, _ = sjson.Delete(mainContent, "children")
	}
	attributes := content.Get("attributes").Raw
	if attributes == "" {
		attributes = "{}"
	}
	separateContent := `{"attributes":` + attributes + `,"children":[` + strings.Join(separateChildren, ",") + `]}`
	return `{"` + className + `":` + mainContent + `}`, `{"` + className + `":` + separateContent + `}`
}

// SeparateRequestDelay is the time PostSeparate waits between the two requests.
// The device applies some parent changes asynchronously, e.g. an Ethernet
// switchport mode change takes about half a second to become operational.
var SeparateRequestDelay = 2 * time.Second

// PostSeparate posts body to dn and sends objects of the given classes in a
// second request after the rest of the body. Some objects are only applied by
// the device if their parent already exists with its final configuration, for
// example VLAN mappings that depend on the switchport mode set in the same
// request. Sent in the same request, or right after the parent change, the
// device accepts them but does not apply them.
func PostSeparate(client *nxos.Client, dn, body string, classNames []string) (nxos.Res, error) {
	mainBody, separateBody := SplitSeparateRequestBody(body, classNames)
	res, err := client.Post(dn, mainBody)
	if err != nil || separateBody == "" {
		return res, err
	}
	time.Sleep(SeparateRequestDelay)
	return client.Post(dn, separateBody)
}
