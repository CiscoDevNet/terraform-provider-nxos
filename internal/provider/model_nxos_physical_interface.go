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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PhysicalInterface struct {
	Device                 types.String `tfsdk:"device"`
	Dn                     types.String `tfsdk:"id"`
	InterfaceId            types.String `tfsdk:"interface_id"`
	FecMode                types.String `tfsdk:"fec_mode"`
	AccessVlan             types.String `tfsdk:"access_vlan"`
	AdminState             types.String `tfsdk:"admin_state"`
	AutoNegotiation        types.String `tfsdk:"auto_negotiation"`
	Bandwidth              types.Int64  `tfsdk:"bandwidth"`
	Delay                  types.Int64  `tfsdk:"delay"`
	Description            types.String `tfsdk:"description"`
	Duplex                 types.String `tfsdk:"duplex"`
	Layer                  types.String `tfsdk:"layer"`
	LinkLogging            types.String `tfsdk:"link_logging"`
	LinkDebounceDown       types.Int64  `tfsdk:"link_debounce_down"`
	LinkDebounceUp         types.Int64  `tfsdk:"link_debounce_up"`
	Medium                 types.String `tfsdk:"medium"`
	Mode                   types.String `tfsdk:"mode"`
	Mtu                    types.Int64  `tfsdk:"mtu"`
	NativeVlan             types.String `tfsdk:"native_vlan"`
	Speed                  types.String `tfsdk:"speed"`
	SpeedGroup             types.String `tfsdk:"speed_group"`
	TrunkVlans             types.String `tfsdk:"trunk_vlans"`
	UniDirectionalEthernet types.String `tfsdk:"uni_directional_ethernet"`
	UserConfiguredFlags    types.String `tfsdk:"user_configured_flags"`
}

func (data PhysicalInterface) getDn() string {
	return fmt.Sprintf("sys/intf/phys-[%s]", data.InterfaceId.ValueString())
}

func (data PhysicalInterface) getClassName() string {
	return "l1PhysIf"
}

func (data PhysicalInterface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.FecMode.IsUnknown() && !data.FecMode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"FECMode", data.FecMode.ValueString())
	}
	if (!data.AccessVlan.IsUnknown() && !data.AccessVlan.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"accessVlan", data.AccessVlan.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.AutoNegotiation.IsUnknown() && !data.AutoNegotiation.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"autoNeg", data.AutoNegotiation.ValueString())
	}
	if (!data.Bandwidth.IsUnknown() && !data.Bandwidth.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bw", strconv.FormatInt(data.Bandwidth.ValueInt64(), 10))
	}
	if (!data.Delay.IsUnknown() && !data.Delay.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"delay", strconv.FormatInt(data.Delay.ValueInt64(), 10))
	}
	if (!data.Description.IsUnknown() && !data.Description.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"descr", data.Description.ValueString())
	}
	if (!data.Duplex.IsUnknown() && !data.Duplex.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"duplex", data.Duplex.ValueString())
	}
	if (!data.Layer.IsUnknown() && !data.Layer.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"layer", data.Layer.ValueString())
	}
	if (!data.LinkLogging.IsUnknown() && !data.LinkLogging.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkLog", data.LinkLogging.ValueString())
	}
	if (!data.LinkDebounceDown.IsUnknown() && !data.LinkDebounceDown.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkDebounce", strconv.FormatInt(data.LinkDebounceDown.ValueInt64(), 10))
	}
	if (!data.LinkDebounceUp.IsUnknown() && !data.LinkDebounceUp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkDebounceLinkUp", strconv.FormatInt(data.LinkDebounceUp.ValueInt64(), 10))
	}
	if (!data.Medium.IsUnknown() && !data.Medium.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"medium", data.Medium.ValueString())
	}
	if (!data.Mode.IsUnknown() && !data.Mode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mode", data.Mode.ValueString())
	}
	if (!data.Mtu.IsUnknown() && !data.Mtu.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if (!data.NativeVlan.IsUnknown() && !data.NativeVlan.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nativeVlan", data.NativeVlan.ValueString())
	}
	if (!data.Speed.IsUnknown() && !data.Speed.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"speed", data.Speed.ValueString())
	}
	if (!data.SpeedGroup.IsUnknown() && !data.SpeedGroup.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"speedGroup", data.SpeedGroup.ValueString())
	}
	if (!data.TrunkVlans.IsUnknown() && !data.TrunkVlans.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"trunkVlans", data.TrunkVlans.ValueString())
	}
	if (!data.UniDirectionalEthernet.IsUnknown() && !data.UniDirectionalEthernet.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"uniDirectionalEthernet", data.UniDirectionalEthernet.ValueString())
	}
	if (!data.UserConfiguredFlags.IsUnknown() && !data.UserConfiguredFlags.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"userCfgdFlags", data.UserConfiguredFlags.ValueString())
	}

	return nxos.Body{body}
}

func (data *PhysicalInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.FecMode.IsNull() || all {
		data.FecMode = types.StringValue(res.Get(data.getClassName() + ".attributes.FECMode").String())
	} else {
		data.FecMode = types.StringNull()
	}
	if !data.AccessVlan.IsNull() || all {
		data.AccessVlan = types.StringValue(res.Get(data.getClassName() + ".attributes.accessVlan").String())
	} else {
		data.AccessVlan = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
	if !data.AutoNegotiation.IsNull() || all {
		data.AutoNegotiation = types.StringValue(res.Get(data.getClassName() + ".attributes.autoNeg").String())
	} else {
		data.AutoNegotiation = types.StringNull()
	}
	if !data.Bandwidth.IsNull() || all {
		data.Bandwidth = types.Int64Value(res.Get(data.getClassName() + ".attributes.bw").Int())
	} else {
		data.Bandwidth = types.Int64Null()
	}
	if !data.Delay.IsNull() || all {
		data.Delay = types.Int64Value(res.Get(data.getClassName() + ".attributes.delay").Int())
	} else {
		data.Delay = types.Int64Null()
	}
	if !data.Description.IsNull() || all {
		data.Description = types.StringValue(res.Get(data.getClassName() + ".attributes.descr").String())
	} else {
		data.Description = types.StringNull()
	}
	if !data.Duplex.IsNull() || all {
		data.Duplex = types.StringValue(res.Get(data.getClassName() + ".attributes.duplex").String())
	} else {
		data.Duplex = types.StringNull()
	}
	if !data.Layer.IsNull() || all {
		data.Layer = types.StringValue(res.Get(data.getClassName() + ".attributes.layer").String())
	} else {
		data.Layer = types.StringNull()
	}
	if !data.LinkLogging.IsNull() || all {
		data.LinkLogging = types.StringValue(res.Get(data.getClassName() + ".attributes.linkLog").String())
	} else {
		data.LinkLogging = types.StringNull()
	}
	if !data.LinkDebounceDown.IsNull() || all {
		data.LinkDebounceDown = types.Int64Value(res.Get(data.getClassName() + ".attributes.linkDebounce").Int())
	} else {
		data.LinkDebounceDown = types.Int64Null()
	}
	if !data.LinkDebounceUp.IsNull() || all {
		data.LinkDebounceUp = types.Int64Value(res.Get(data.getClassName() + ".attributes.linkDebounceLinkUp").Int())
	} else {
		data.LinkDebounceUp = types.Int64Null()
	}
	if !data.Medium.IsNull() || all {
		data.Medium = types.StringValue(res.Get(data.getClassName() + ".attributes.medium").String())
	} else {
		data.Medium = types.StringNull()
	}
	if !data.Mode.IsNull() || all {
		data.Mode = types.StringValue(res.Get(data.getClassName() + ".attributes.mode").String())
	} else {
		data.Mode = types.StringNull()
	}
	if !data.Mtu.IsNull() || all {
		data.Mtu = types.Int64Value(res.Get(data.getClassName() + ".attributes.mtu").Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if !data.NativeVlan.IsNull() || all {
		data.NativeVlan = types.StringValue(res.Get(data.getClassName() + ".attributes.nativeVlan").String())
	} else {
		data.NativeVlan = types.StringNull()
	}
	if !data.Speed.IsNull() || all {
		data.Speed = types.StringValue(res.Get(data.getClassName() + ".attributes.speed").String())
	} else {
		data.Speed = types.StringNull()
	}
	if !data.SpeedGroup.IsNull() || all {
		data.SpeedGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.speedGroup").String())
	} else {
		data.SpeedGroup = types.StringNull()
	}
	if !data.TrunkVlans.IsNull() || all {
		data.TrunkVlans = types.StringValue(res.Get(data.getClassName() + ".attributes.trunkVlans").String())
	} else {
		data.TrunkVlans = types.StringNull()
	}
	if !data.UniDirectionalEthernet.IsNull() || all {
		data.UniDirectionalEthernet = types.StringValue(res.Get(data.getClassName() + ".attributes.uniDirectionalEthernet").String())
	} else {
		data.UniDirectionalEthernet = types.StringNull()
	}
	if !data.UserConfiguredFlags.IsNull() || all {
		data.UserConfiguredFlags = types.StringValue(res.Get(data.getClassName() + ".attributes.userCfgdFlags").String())
	} else {
		data.UserConfiguredFlags = types.StringNull()
	}
}

func (data PhysicalInterface) toDeleteBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"FECMode", "auto")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"accessVlan", "vlan-1")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", "DME_UNSET_PROPERTY_MARKER")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"autoNeg", "on")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bw", strconv.FormatInt(0, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"delay", strconv.FormatInt(1, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"descr", "DME_UNSET_PROPERTY_MARKER")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"duplex", "auto")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"layer", "Layer2")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkLog", "default")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkDebounce", strconv.FormatInt(100, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkDebounceLinkUp", strconv.FormatInt(0, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"medium", "broadcast")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mode", "access")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtu", strconv.FormatInt(1500, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nativeVlan", "vlan-1")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"speed", "auto")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"speedGroup", "auto")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"trunkVlans", "1-4094")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"uniDirectionalEthernet", "disable")
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"userCfgdFlags", "none")

	return nxos.Body{body}
}

func (data *PhysicalInterface) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/intf/phys-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InterfaceId = types.StringValue(matches[1])
}
