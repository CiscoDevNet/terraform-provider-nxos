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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type EVPNVNIRouteTarget struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	Encap       types.String `tfsdk:"encap"`
	Direction   types.String `tfsdk:"direction"`
	RouteTarget types.String `tfsdk:"route_target"`
}

func (data EVPNVNIRouteTarget) getDn() string {
	return fmt.Sprintf("sys/evpn/bdevi-[%s]/rttp-[%s]/ent-[%s]", data.Encap.ValueString(), data.Direction.ValueString(), data.RouteTarget.ValueString())
}

func (data EVPNVNIRouteTarget) getClassName() string {
	return "rtctrlRttEntry"
}

func (data EVPNVNIRouteTarget) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.RouteTarget.IsUnknown() && !data.RouteTarget.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rtt", data.RouteTarget.ValueString())
	}

	return nxos.Body{body}
}

func (data *EVPNVNIRouteTarget) fromBody(res gjson.Result, all bool) {
	if !data.RouteTarget.IsNull() || all {
		data.RouteTarget = types.StringValue(res.Get(data.getClassName() + ".attributes.rtt").String())
	} else {
		data.RouteTarget = types.StringNull()
	}
}

func (data EVPNVNIRouteTarget) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *EVPNVNIRouteTarget) getIdsFromDn() {
	reString := "sys/evpn/bdevi-[%s]/rttp-[%s]/ent-[%s]"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Encap = types.StringValue(matches[1])
	data.Direction = types.StringValue(matches[2])
	data.RouteTarget = types.StringValue(matches[3])
}
