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

type ICMPv4Interface struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	VrfName     types.String `tfsdk:"vrf_name"`
	InterfaceId types.String `tfsdk:"interface_id"`
	Control     types.String `tfsdk:"control"`
}

func (data ICMPv4Interface) getDn() string {
	return fmt.Sprintf("sys/icmpv4/inst/dom-[%s]/if-[%s]", data.VrfName.ValueString(), data.InterfaceId.ValueString())
}

func (data ICMPv4Interface) getClassName() string {
	return "icmpv4If"
}

func (data ICMPv4Interface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.Control.IsUnknown() && !data.Control.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ctrl", data.Control.ValueString())
	}

	return nxos.Body{body}
}

func (data *ICMPv4Interface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.Control.IsNull() || all {
		data.Control = types.StringValue(res.Get(data.getClassName() + ".attributes.ctrl").String())
	} else {
		data.Control = types.StringNull()
	}
}

func (data ICMPv4Interface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *ICMPv4Interface) getIdsFromDn() {
	reString := "sys/icmpv4/inst/dom-[%s]/if-[%s]"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.VrfName = types.StringValue(matches[1])
	data.InterfaceId = types.StringValue(matches[2])
}
