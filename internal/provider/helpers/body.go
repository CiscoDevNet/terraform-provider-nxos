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
