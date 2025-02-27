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

type IPv4Interface struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	Vrf         types.String `tfsdk:"vrf"`
	InterfaceId types.String `tfsdk:"interface_id"`
	DropGlean   types.String `tfsdk:"drop_glean"`
	Forward     types.String `tfsdk:"forward"`
	Unnumbered  types.String `tfsdk:"unnumbered"`
	Urpf        types.String `tfsdk:"urpf"`
}

func (data IPv4Interface) getDn() string {
	return fmt.Sprintf("sys/ipv4/inst/dom-[%s]/if-[%s]", data.Vrf.ValueString(), data.InterfaceId.ValueString())
}

func (data IPv4Interface) getClassName() string {
	return "ipv4If"
}

func (data IPv4Interface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.DropGlean.IsUnknown() && !data.DropGlean.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dropGlean", data.DropGlean.ValueString())
	}
	if (!data.Forward.IsUnknown() && !data.Forward.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"forward", data.Forward.ValueString())
	}
	if (!data.Unnumbered.IsUnknown() && !data.Unnumbered.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"unnumbered", data.Unnumbered.ValueString())
	}
	if (!data.Urpf.IsUnknown() && !data.Urpf.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"urpf", data.Urpf.ValueString())
	}

	return nxos.Body{body}
}

func (data *IPv4Interface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.DropGlean.IsNull() || all {
		data.DropGlean = types.StringValue(res.Get(data.getClassName() + ".attributes.dropGlean").String())
	} else {
		data.DropGlean = types.StringNull()
	}
	if !data.Forward.IsNull() || all {
		data.Forward = types.StringValue(res.Get(data.getClassName() + ".attributes.forward").String())
	} else {
		data.Forward = types.StringNull()
	}
	if !data.Unnumbered.IsNull() || all {
		data.Unnumbered = types.StringValue(res.Get(data.getClassName() + ".attributes.unnumbered").String())
	} else {
		data.Unnumbered = types.StringNull()
	}
	if !data.Urpf.IsNull() || all {
		data.Urpf = types.StringValue(res.Get(data.getClassName() + ".attributes.urpf").String())
	} else {
		data.Urpf = types.StringNull()
	}
}

func (data IPv4Interface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *IPv4Interface) getIdsFromDn() {
	reString := "sys/ipv4/inst/dom-[%s]/if-[%s]"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.InterfaceId = types.StringValue(matches[2])
}
