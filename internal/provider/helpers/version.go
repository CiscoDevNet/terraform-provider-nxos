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
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

// MinVersion is the minimum NX-OS version of a class or attribute, together with
// the Terraform attribute path used in error messages.
type MinVersion struct {
	Version string
	Path    string
}

// SubtreeClasses returns the comma-separated list of classes to be used as `rsp-subtree-class` query parameter.
// Classes with a minimum version higher than the device version are omitted, as NX-OS rejects queries
// containing unknown classes. If the device version cannot be determined, all classes are returned.
// The device version is only retrieved if minVersions is not empty.
func SubtreeClasses(ctx context.Context, client *nxos.Client, classes []string, minVersions map[string]string) string {
	if len(minVersions) == 0 {
		return strings.Join(classes, ",")
	}
	version, err := client.Version()
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to retrieve NX-OS version, querying all classes: %s", err))
		return strings.Join(classes, ",")
	}
	var result []string
	for _, c := range classes {
		if v, ok := minVersions[c]; ok && !version.AtLeast(nxos.MustParseVersion(v)) {
			continue
		}
		result = append(result, c)
	}
	return strings.Join(result, ",")
}

// CheckMinVersion verifies that a request body does not contain classes or attributes that are not supported
// by the device version. checks is keyed by class name, or "<class name>.<attribute name>" for attributes.
// Objects marked for deletion and properties being unset are ignored. If the device version cannot be determined,
// no error is returned and the device will validate the request. The device version is only retrieved
// if checks is not empty.
func CheckMinVersion(ctx context.Context, client *nxos.Client, device, body string, checks map[string]MinVersion) diag.Diagnostics {
	var diags diag.Diagnostics
	if len(checks) == 0 {
		return diags
	}
	version, err := client.Version()
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Failed to retrieve NX-OS version, skipping minimum version check: %s", err))
		return diags
	}
	violations := map[string]string{}
	var walk func(gjson.Result)
	walk = func(node gjson.Result) {
		node.ForEach(func(className, obj gjson.Result) bool {
			attrs := obj.Get("attributes")
			if attrs.Get("status").String() == "deleted" {
				return true
			}
			if c, ok := checks[className.String()]; ok && !version.AtLeast(nxos.MustParseVersion(c.Version)) {
				violations[c.Path] = c.Version
			}
			attrs.ForEach(func(name, value gjson.Result) bool {
				if value.String() == "DME_UNSET_PROPERTY_MARKER" {
					return true
				}
				if c, ok := checks[className.String()+"."+name.String()]; ok && !version.AtLeast(nxos.MustParseVersion(c.Version)) {
					violations[c.Path] = c.Version
				}
				return true
			})
			for _, child := range obj.Get("children").Array() {
				walk(child)
			}
			return true
		})
	}
	walk(gjson.Parse(body))

	paths := make([]string, 0, len(violations))
	for p := range violations {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	deviceName := ""
	if device != "" {
		deviceName = fmt.Sprintf(" '%s'", device)
	}
	for _, p := range paths {
		diags.AddError(
			"Unsupported NX-OS Version",
			fmt.Sprintf("Attribute `%s` requires NX-OS %s or later, but device%s runs NX-OS %s.", p, violations[p], deviceName, version),
		)
	}
	return diags
}
