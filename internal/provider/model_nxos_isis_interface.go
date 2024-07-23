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

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ISISInterface struct {
	Device                types.String `tfsdk:"device"`
	Dn                    types.String `tfsdk:"id"`
	InterfaceId           types.String `tfsdk:"interface_id"`
	AuthenticationCheck   types.Bool   `tfsdk:"authentication_check"`
	AuthenticationCheckL1 types.Bool   `tfsdk:"authentication_check_l1"`
	AuthenticationCheckL2 types.Bool   `tfsdk:"authentication_check_l2"`
	AuthenticationKey     types.String `tfsdk:"authentication_key"`
	AuthenticationKeyL1   types.String `tfsdk:"authentication_key_l1"`
	AuthenticationKeyL2   types.String `tfsdk:"authentication_key_l2"`
	AuthenticationType    types.String `tfsdk:"authentication_type"`
	AuthenticationTypeL1  types.String `tfsdk:"authentication_type_l1"`
	AuthenticationTypeL2  types.String `tfsdk:"authentication_type_l2"`
	CircuitType           types.String `tfsdk:"circuit_type"`
	Vrf                   types.String `tfsdk:"vrf"`
	HelloInterval         types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalL1       types.Int64  `tfsdk:"hello_interval_l1"`
	HelloIntervalL2       types.Int64  `tfsdk:"hello_interval_l2"`
	HelloMultiplier       types.Int64  `tfsdk:"hello_multiplier"`
	HelloMultiplierL1     types.Int64  `tfsdk:"hello_multiplier_l1"`
	HelloMultiplierL2     types.Int64  `tfsdk:"hello_multiplier_l2"`
	HelloPadding          types.String `tfsdk:"hello_padding"`
	InstanceName          types.String `tfsdk:"instance_name"`
	MetricL1              types.Int64  `tfsdk:"metric_l1"`
	MetricL2              types.Int64  `tfsdk:"metric_l2"`
	MtuCheck              types.Bool   `tfsdk:"mtu_check"`
	MtuCheckL1            types.Bool   `tfsdk:"mtu_check_l1"`
	MtuCheckL2            types.Bool   `tfsdk:"mtu_check_l2"`
	NetworkTypeP2p        types.String `tfsdk:"network_type_p2p"`
	Passive               types.String `tfsdk:"passive"`
	PriorityL1            types.Int64  `tfsdk:"priority_l1"`
	PriorityL2            types.Int64  `tfsdk:"priority_l2"`
	EnableIpv4            types.Bool   `tfsdk:"enable_ipv4"`
}

func (data ISISInterface) getDn() string {
	return fmt.Sprintf("sys/isis/if-[%s]", data.InterfaceId.ValueString())
}

func (data ISISInterface) getClassName() string {
	return "isisInternalIf"
}

func (data ISISInterface) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.AuthenticationCheck.IsUnknown() && !data.AuthenticationCheck.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authCheck", strconv.FormatBool(data.AuthenticationCheck.ValueBool()))
	}
	if (!data.AuthenticationCheckL1.IsUnknown() && !data.AuthenticationCheckL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authCheckLvl1", strconv.FormatBool(data.AuthenticationCheckL1.ValueBool()))
	}
	if (!data.AuthenticationCheckL2.IsUnknown() && !data.AuthenticationCheckL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authCheckLvl2", strconv.FormatBool(data.AuthenticationCheckL2.ValueBool()))
	}
	if (!data.AuthenticationKey.IsUnknown() && !data.AuthenticationKey.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authKey", data.AuthenticationKey.ValueString())
	}
	if (!data.AuthenticationKeyL1.IsUnknown() && !data.AuthenticationKeyL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authKeyLvl1", data.AuthenticationKeyL1.ValueString())
	}
	if (!data.AuthenticationKeyL2.IsUnknown() && !data.AuthenticationKeyL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authKeyLvl2", data.AuthenticationKeyL2.ValueString())
	}
	if (!data.AuthenticationType.IsUnknown() && !data.AuthenticationType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authType", data.AuthenticationType.ValueString())
	}
	if (!data.AuthenticationTypeL1.IsUnknown() && !data.AuthenticationTypeL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authTypeLvl1", data.AuthenticationTypeL1.ValueString())
	}
	if (!data.AuthenticationTypeL2.IsUnknown() && !data.AuthenticationTypeL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authTypeLvl2", data.AuthenticationTypeL2.ValueString())
	}
	if (!data.CircuitType.IsUnknown() && !data.CircuitType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cktT", data.CircuitType.ValueString())
	}
	if (!data.Vrf.IsUnknown() && !data.Vrf.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dom", data.Vrf.ValueString())
	}
	if (!data.HelloInterval.IsUnknown() && !data.HelloInterval.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloIntvl", strconv.FormatInt(data.HelloInterval.ValueInt64(), 10))
	}
	if (!data.HelloIntervalL1.IsUnknown() && !data.HelloIntervalL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloIntvlLvl1", strconv.FormatInt(data.HelloIntervalL1.ValueInt64(), 10))
	}
	if (!data.HelloIntervalL2.IsUnknown() && !data.HelloIntervalL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloIntvlLvl2", strconv.FormatInt(data.HelloIntervalL2.ValueInt64(), 10))
	}
	if (!data.HelloMultiplier.IsUnknown() && !data.HelloMultiplier.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloMult", strconv.FormatInt(data.HelloMultiplier.ValueInt64(), 10))
	}
	if (!data.HelloMultiplierL1.IsUnknown() && !data.HelloMultiplierL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloMultLvl1", strconv.FormatInt(data.HelloMultiplierL1.ValueInt64(), 10))
	}
	if (!data.HelloMultiplierL2.IsUnknown() && !data.HelloMultiplierL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloMultLvl2", strconv.FormatInt(data.HelloMultiplierL2.ValueInt64(), 10))
	}
	if (!data.HelloPadding.IsUnknown() && !data.HelloPadding.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloPad", data.HelloPadding.ValueString())
	}
	if (!data.InstanceName.IsUnknown() && !data.InstanceName.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"instance", data.InstanceName.ValueString())
	}
	if (!data.MetricL1.IsUnknown() && !data.MetricL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"metricLvl1", strconv.FormatInt(data.MetricL1.ValueInt64(), 10))
	}
	if (!data.MetricL2.IsUnknown() && !data.MetricL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"metricLvl2", strconv.FormatInt(data.MetricL2.ValueInt64(), 10))
	}
	if (!data.MtuCheck.IsUnknown() && !data.MtuCheck.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtuCheck", strconv.FormatBool(data.MtuCheck.ValueBool()))
	}
	if (!data.MtuCheckL1.IsUnknown() && !data.MtuCheckL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtuCheckLvl1", strconv.FormatBool(data.MtuCheckL1.ValueBool()))
	}
	if (!data.MtuCheckL2.IsUnknown() && !data.MtuCheckL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtuCheckLvl2", strconv.FormatBool(data.MtuCheckL2.ValueBool()))
	}
	if (!data.NetworkTypeP2p.IsUnknown() && !data.NetworkTypeP2p.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"networkTypeP2P", data.NetworkTypeP2p.ValueString())
	}
	if (!data.Passive.IsUnknown() && !data.Passive.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"passive", data.Passive.ValueString())
	}
	if (!data.PriorityL1.IsUnknown() && !data.PriorityL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"priorityLvl1", strconv.FormatInt(data.PriorityL1.ValueInt64(), 10))
	}
	if (!data.PriorityL2.IsUnknown() && !data.PriorityL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"priorityLvl2", strconv.FormatInt(data.PriorityL2.ValueInt64(), 10))
	}
	if (!data.EnableIpv4.IsUnknown() && !data.EnableIpv4.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"v4enable", strconv.FormatBool(data.EnableIpv4.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *ISISInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.AuthenticationCheck.IsNull() || all {
		data.AuthenticationCheck = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.authCheck").String()))
	} else {
		data.AuthenticationCheck = types.BoolNull()
	}
	if !data.AuthenticationCheckL1.IsNull() || all {
		data.AuthenticationCheckL1 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.authCheckLvl1").String()))
	} else {
		data.AuthenticationCheckL1 = types.BoolNull()
	}
	if !data.AuthenticationCheckL2.IsNull() || all {
		data.AuthenticationCheckL2 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.authCheckLvl2").String()))
	} else {
		data.AuthenticationCheckL2 = types.BoolNull()
	}
	if !data.AuthenticationType.IsNull() || all {
		data.AuthenticationType = types.StringValue(res.Get(data.getClassName() + ".attributes.authType").String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if !data.AuthenticationTypeL1.IsNull() || all {
		data.AuthenticationTypeL1 = types.StringValue(res.Get(data.getClassName() + ".attributes.authTypeLvl1").String())
	} else {
		data.AuthenticationTypeL1 = types.StringNull()
	}
	if !data.AuthenticationTypeL2.IsNull() || all {
		data.AuthenticationTypeL2 = types.StringValue(res.Get(data.getClassName() + ".attributes.authTypeLvl2").String())
	} else {
		data.AuthenticationTypeL2 = types.StringNull()
	}
	if !data.CircuitType.IsNull() || all {
		data.CircuitType = types.StringValue(res.Get(data.getClassName() + ".attributes.cktT").String())
	} else {
		data.CircuitType = types.StringNull()
	}
	if !data.Vrf.IsNull() || all {
		data.Vrf = types.StringValue(res.Get(data.getClassName() + ".attributes.dom").String())
	} else {
		data.Vrf = types.StringNull()
	}
	if !data.HelloInterval.IsNull() || all {
		data.HelloInterval = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloIntvl").Int())
	} else {
		data.HelloInterval = types.Int64Null()
	}
	if !data.HelloIntervalL1.IsNull() || all {
		data.HelloIntervalL1 = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloIntvlLvl1").Int())
	} else {
		data.HelloIntervalL1 = types.Int64Null()
	}
	if !data.HelloIntervalL2.IsNull() || all {
		data.HelloIntervalL2 = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloIntvlLvl2").Int())
	} else {
		data.HelloIntervalL2 = types.Int64Null()
	}
	if !data.HelloMultiplier.IsNull() || all {
		data.HelloMultiplier = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloMult").Int())
	} else {
		data.HelloMultiplier = types.Int64Null()
	}
	if !data.HelloMultiplierL1.IsNull() || all {
		data.HelloMultiplierL1 = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloMultLvl1").Int())
	} else {
		data.HelloMultiplierL1 = types.Int64Null()
	}
	if !data.HelloMultiplierL2.IsNull() || all {
		data.HelloMultiplierL2 = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloMultLvl2").Int())
	} else {
		data.HelloMultiplierL2 = types.Int64Null()
	}
	if !data.HelloPadding.IsNull() || all {
		data.HelloPadding = types.StringValue(res.Get(data.getClassName() + ".attributes.helloPad").String())
	} else {
		data.HelloPadding = types.StringNull()
	}
	if !data.InstanceName.IsNull() || all {
		data.InstanceName = types.StringValue(res.Get(data.getClassName() + ".attributes.instance").String())
	} else {
		data.InstanceName = types.StringNull()
	}
	if !data.MetricL1.IsNull() || all {
		data.MetricL1 = types.Int64Value(res.Get(data.getClassName() + ".attributes.metricLvl1").Int())
	} else {
		data.MetricL1 = types.Int64Null()
	}
	if !data.MetricL2.IsNull() || all {
		data.MetricL2 = types.Int64Value(res.Get(data.getClassName() + ".attributes.metricLvl2").Int())
	} else {
		data.MetricL2 = types.Int64Null()
	}
	if !data.MtuCheck.IsNull() || all {
		data.MtuCheck = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.mtuCheck").String()))
	} else {
		data.MtuCheck = types.BoolNull()
	}
	if !data.MtuCheckL1.IsNull() || all {
		data.MtuCheckL1 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.mtuCheckLvl1").String()))
	} else {
		data.MtuCheckL1 = types.BoolNull()
	}
	if !data.MtuCheckL2.IsNull() || all {
		data.MtuCheckL2 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.mtuCheckLvl2").String()))
	} else {
		data.MtuCheckL2 = types.BoolNull()
	}
	if !data.NetworkTypeP2p.IsNull() || all {
		data.NetworkTypeP2p = types.StringValue(res.Get(data.getClassName() + ".attributes.networkTypeP2P").String())
	} else {
		data.NetworkTypeP2p = types.StringNull()
	}
	if !data.Passive.IsNull() || all {
		data.Passive = types.StringValue(res.Get(data.getClassName() + ".attributes.passive").String())
	} else {
		data.Passive = types.StringNull()
	}
	if !data.PriorityL1.IsNull() || all {
		data.PriorityL1 = types.Int64Value(res.Get(data.getClassName() + ".attributes.priorityLvl1").Int())
	} else {
		data.PriorityL1 = types.Int64Null()
	}
	if !data.PriorityL2.IsNull() || all {
		data.PriorityL2 = types.Int64Value(res.Get(data.getClassName() + ".attributes.priorityLvl2").Int())
	} else {
		data.PriorityL2 = types.Int64Null()
	}
	if !data.EnableIpv4.IsNull() || all {
		data.EnableIpv4 = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.v4enable").String()))
	} else {
		data.EnableIpv4 = types.BoolNull()
	}
}

func (data ISISInterface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
