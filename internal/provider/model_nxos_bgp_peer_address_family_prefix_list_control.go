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

type BGPPeerAddressFamilyPrefixListControl struct {
	Device        types.String `tfsdk:"device"`
	Dn            types.String `tfsdk:"id"`
	Asn           types.String `tfsdk:"asn"`
	Vrf           types.String `tfsdk:"vrf"`
	Address       types.String `tfsdk:"address"`
	AddressFamily types.String `tfsdk:"address_family"`
	Direction     types.String `tfsdk:"direction"`
	List          types.String `tfsdk:"list"`
}

func (data BGPPeerAddressFamilyPrefixListControl) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]/af-[%s]/pfxctrl-[%s]", data.Vrf.ValueString(), data.Address.ValueString(), data.AddressFamily.ValueString(), data.Direction.ValueString())
}

func (data BGPPeerAddressFamilyPrefixListControl) getClassName() string {
	return "bgpPfxCtrlP"
}

func (data BGPPeerAddressFamilyPrefixListControl) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Direction.IsUnknown() && !data.Direction.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"direction", data.Direction.ValueString())
	}
	if (!data.List.IsUnknown() && !data.List.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"list", data.List.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPPeerAddressFamilyPrefixListControl) fromBody(res gjson.Result, all bool) {
	if !data.Direction.IsNull() || all {
		data.Direction = types.StringValue(res.Get(data.getClassName() + ".attributes.direction").String())
	} else {
		data.Direction = types.StringNull()
	}
	if !data.List.IsNull() || all {
		data.List = types.StringValue(res.Get(data.getClassName() + ".attributes.list").String())
	} else {
		data.List = types.StringNull()
	}
}

func (data BGPPeerAddressFamilyPrefixListControl) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *BGPPeerAddressFamilyPrefixListControl) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/bgp/inst/dom-[%s]/peer-[%s]/af-[%s]/pfxctrl-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.Address = types.StringValue(matches[2])
	data.AddressFamily = types.StringValue(matches[3])
	data.Direction = types.StringValue(matches[4])
}
