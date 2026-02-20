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

// ParseImportID splits an import identifier into a DN and an optional device name.
// The format is "<dn>" or "<dn>,<device>". Since DNs can contain commas inside
// brackets (e.g. "sys/acl/ipv4/name-[a,b]"), only the last comma that is not
// enclosed in brackets is treated as the separator.
func ParseImportID(id string) (dn string, device string) {
	// Find the last comma at bracket depth 0
	depth := 0
	sepIdx := -1
	for i, c := range id {
		switch c {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				sepIdx = i
			}
		}
	}
	if sepIdx == -1 {
		return id, ""
	}
	return id[:sepIdx], id[sepIdx+1:]
}

func ParseNxosBoolean(s string) bool {
	if s == "yes" || s == "true" {
		return true
	}
	return false
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
