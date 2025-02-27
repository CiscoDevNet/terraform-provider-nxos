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
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ISISAddressFamily struct {
	Device                   types.String `tfsdk:"device"`
	Dn                       types.String `tfsdk:"id"`
	InstanceName             types.String `tfsdk:"instance_name"`
	Vrf                      types.String `tfsdk:"vrf"`
	AddressFamily            types.String `tfsdk:"address_family"`
	SegmentRoutingMpls       types.Bool   `tfsdk:"segment_routing_mpls"`
	EnableBfd                types.Bool   `tfsdk:"enable_bfd"`
	PrefixAdvertisePassiveL1 types.Bool   `tfsdk:"prefix_advertise_passive_l1"`
	PrefixAdvertisePassiveL2 types.Bool   `tfsdk:"prefix_advertise_passive_l2"`
}

func (data ISISAddressFamily) getDn() string {
	return fmt.Sprintf("sys/isis/inst-[%s]/dom-[%s]/af-[%s]", data.InstanceName.ValueString(), data.Vrf.ValueString(), data.AddressFamily.ValueString())
}

func (data ISISAddressFamily) getClassName() string {
	return "isisDomAf"
}

func (data ISISAddressFamily) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Vrf.IsUnknown() && !data.Vrf.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Vrf.ValueString())
	}
	if (!data.AddressFamily.IsUnknown() && !data.AddressFamily.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.AddressFamily.ValueString())
	}
	if (!data.SegmentRoutingMpls.IsUnknown() && !data.SegmentRoutingMpls.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srMpls", strconv.FormatBool(data.SegmentRoutingMpls.ValueBool()))
	}
	if (!data.EnableBfd.IsUnknown() && !data.EnableBfd.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"enableBfd", strconv.FormatBool(data.EnableBfd.ValueBool()))
	}
	if (!data.PrefixAdvertisePassiveL1.IsUnknown() && !data.PrefixAdvertisePassiveL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advPassiveLvl1", strconv.FormatBool(data.PrefixAdvertisePassiveL1.ValueBool()))
	}
	if (!data.PrefixAdvertisePassiveL2.IsUnknown() && !data.PrefixAdvertisePassiveL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advPassiveLvl2", strconv.FormatBool(data.PrefixAdvertisePassiveL2.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *ISISAddressFamily) fromBody(res gjson.Result, all bool) {
	if !data.Vrf.IsNull() || all {
		data.Vrf = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Vrf = types.StringNull()
	}
	if !data.AddressFamily.IsNull() || all {
		data.AddressFamily = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.AddressFamily = types.StringNull()
	}
	if !data.SegmentRoutingMpls.IsNull() || all {
		data.SegmentRoutingMpls = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.srMpls").String()))
	} else {
		data.SegmentRoutingMpls = types.BoolNull()
	}
	if !data.EnableBfd.IsNull() || all {
		data.EnableBfd = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.enableBfd").String()))
	} else {
		data.EnableBfd = types.BoolNull()
	}
	if !data.PrefixAdvertisePassiveL1.IsNull() || all {
		data.PrefixAdvertisePassiveL1 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.advPassiveLvl1").String()))
	} else {
		data.PrefixAdvertisePassiveL1 = types.BoolNull()
	}
	if !data.PrefixAdvertisePassiveL2.IsNull() || all {
		data.PrefixAdvertisePassiveL2 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.advPassiveLvl2").String()))
	} else {
		data.PrefixAdvertisePassiveL2 = types.BoolNull()
	}
}

func (data ISISAddressFamily) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *ISISAddressFamily) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/isis/inst-[%s]/dom-[%s]/af-[%s]", "%[1]s", "")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InstanceName = types.StringValue(matches[1])
	data.Vrf = types.StringValue(matches[2])
	data.AddressFamily = types.StringValue(matches[3])
}
