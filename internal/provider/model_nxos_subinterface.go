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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Subinterface struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	InterfaceId types.String `tfsdk:"interface_id"`
	AdminState  types.String `tfsdk:"admin_state"`
	Bandwidth   types.Int64  `tfsdk:"bandwidth"`
	Delay       types.Int64  `tfsdk:"delay"`
	Description types.String `tfsdk:"description"`
	Encap       types.String `tfsdk:"encap"`
	LinkLogging types.String `tfsdk:"link_logging"`
	Medium      types.String `tfsdk:"medium"`
	Mtu         types.Int64  `tfsdk:"mtu"`
}

func (data Subinterface) getDn() string {
	return fmt.Sprintf("sys/intf/encrtd-[%s]", data.InterfaceId.ValueString())
}

func (data Subinterface) getClassName() string {
	return "l3EncRtdIf"
}

func (data Subinterface) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
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
	if (!data.Encap.IsUnknown() && !data.Encap.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"encap", data.Encap.ValueString())
	}
	if (!data.LinkLogging.IsUnknown() && !data.LinkLogging.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"linkLog", data.LinkLogging.ValueString())
	}
	if (!data.Medium.IsUnknown() && !data.Medium.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mediumType", data.Medium.ValueString())
	}
	if (!data.Mtu.IsUnknown() && !data.Mtu.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *Subinterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
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
	if !data.Encap.IsNull() || all {
		data.Encap = types.StringValue(res.Get(data.getClassName() + ".attributes.encap").String())
	} else {
		data.Encap = types.StringNull()
	}
	if !data.LinkLogging.IsNull() || all {
		data.LinkLogging = types.StringValue(res.Get(data.getClassName() + ".attributes.linkLog").String())
	} else {
		data.LinkLogging = types.StringNull()
	}
	if !data.Medium.IsNull() || all {
		data.Medium = types.StringValue(res.Get(data.getClassName() + ".attributes.mediumType").String())
	} else {
		data.Medium = types.StringNull()
	}
	if !data.Mtu.IsNull() || all {
		data.Mtu = types.Int64Value(res.Get(data.getClassName() + ".attributes.mtu").Int())
	} else {
		data.Mtu = types.Int64Null()
	}
}

func (data Subinterface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
