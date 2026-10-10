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

package provider

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// GetSubtree reads the object at dn including its subtree, restricted to the given child classes.
// Classes with a minimum version higher than the device version (minVersions) are omitted. On a best
// effort basis, classes the device reports as unknown (e.g. on untested NX-OS releases) are omitted as
// well and remembered for subsequent reads, so that reads do not fail on releases lacking some of them.
func (d *NxosProviderDataDevice) GetSubtree(ctx context.Context, dn string, classes []string, minVersions map[string]string) (nxos.Res, error) {
	classes = slices.DeleteFunc(slices.Clone(helpers.SubtreeClasses(ctx, d.Client, classes, minVersions)), d.isUnsupportedClass)
	for {
		queries := []func(*nxos.Req){nxos.Query("rsp-subtree", "no")}
		if len(classes) > 0 {
			queries = []func(*nxos.Req){nxos.Query("rsp-subtree", "full"), nxos.Query("rsp-subtree-class", strings.Join(classes, ","))}
		}
		res, err := d.Client.GetDn(dn, queries...)
		if err == nil {
			return res, nil
		}
		// NX-OS only reports the first unknown class, e.g. {"imdata":[{"error":{"attributes":{"code":"17","text":"Unknown class esgMatchVlanInterface"}}}]}
		class, ok := strings.CutPrefix(res.Get("imdata.0.error.attributes.text").String(), "Unknown class ")
		if res.Get("imdata.0.error.attributes.code").String() != "17" || !ok || !slices.Contains(classes, class) {
			return res, err
		}
		d.setUnsupportedClass(ctx, class)
		classes = slices.DeleteFunc(classes, func(c string) bool { return c == class })
	}
}

func (d *NxosProviderDataDevice) isUnsupportedClass(class string) bool {
	d.unsupportedClassesMutex.Lock()
	defer d.unsupportedClassesMutex.Unlock()
	return d.unsupportedClasses[class]
}

func (d *NxosProviderDataDevice) setUnsupportedClass(ctx context.Context, class string) {
	d.unsupportedClassesMutex.Lock()
	defer d.unsupportedClassesMutex.Unlock()
	if d.unsupportedClasses == nil {
		d.unsupportedClasses = map[string]bool{}
	}
	if !d.unsupportedClasses[class] {
		tflog.Warn(ctx, fmt.Sprintf("Class %s is not supported by the device %s, ignoring it", class, d.Client.Url))
	}
	d.unsupportedClasses[class] = true
}
