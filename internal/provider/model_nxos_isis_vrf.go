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

type ISISVRF struct {
	Device                types.String `tfsdk:"device"`
	Dn                    types.String `tfsdk:"id"`
	InstanceName          types.String `tfsdk:"instance_name"`
	Name                  types.String `tfsdk:"name"`
	AdminState            types.String `tfsdk:"admin_state"`
	AuthenticationCheckL1 types.Bool   `tfsdk:"authentication_check_l1"`
	AuthenticationCheckL2 types.Bool   `tfsdk:"authentication_check_l2"`
	AuthenticationKeyL1   types.String `tfsdk:"authentication_key_l1"`
	AuthenticationKeyL2   types.String `tfsdk:"authentication_key_l2"`
	AuthenticationTypeL1  types.String `tfsdk:"authentication_type_l1"`
	AuthenticationTypeL2  types.String `tfsdk:"authentication_type_l2"`
	BandwidthReference    types.Int64  `tfsdk:"bandwidth_reference"`
	BanwidthReferenceUnit types.String `tfsdk:"banwidth_reference_unit"`
	IsType                types.String `tfsdk:"is_type"`
	MetricType            types.String `tfsdk:"metric_type"`
	Mtu                   types.Int64  `tfsdk:"mtu"`
	Net                   types.String `tfsdk:"net"`
	PassiveDefault        types.String `tfsdk:"passive_default"`
}

func (data ISISVRF) getDn() string {
	return fmt.Sprintf("sys/isis/inst-[%s]/dom-[%s]", data.InstanceName.ValueString(), data.Name.ValueString())
}

func (data ISISVRF) getClassName() string {
	return "isisDom"
}

func (data ISISVRF) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.AuthenticationCheckL1.IsUnknown() && !data.AuthenticationCheckL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authCheckLvl1", strconv.FormatBool(data.AuthenticationCheckL1.ValueBool()))
	}
	if (!data.AuthenticationCheckL2.IsUnknown() && !data.AuthenticationCheckL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authCheckLvl2", strconv.FormatBool(data.AuthenticationCheckL2.ValueBool()))
	}
	if (!data.AuthenticationKeyL1.IsUnknown() && !data.AuthenticationKeyL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authKeyLvl1", data.AuthenticationKeyL1.ValueString())
	}
	if (!data.AuthenticationKeyL2.IsUnknown() && !data.AuthenticationKeyL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authKeyLvl2", data.AuthenticationKeyL2.ValueString())
	}
	if (!data.AuthenticationTypeL1.IsUnknown() && !data.AuthenticationTypeL1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authTypeLvl1", data.AuthenticationTypeL1.ValueString())
	}
	if (!data.AuthenticationTypeL2.IsUnknown() && !data.AuthenticationTypeL2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authTypeLvl2", data.AuthenticationTypeL2.ValueString())
	}
	if (!data.BandwidthReference.IsUnknown() && !data.BandwidthReference.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bwRef", strconv.FormatInt(data.BandwidthReference.ValueInt64(), 10))
	}
	if (!data.BanwidthReferenceUnit.IsUnknown() && !data.BanwidthReferenceUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bwRefUnit", data.BanwidthReferenceUnit.ValueString())
	}
	if (!data.IsType.IsUnknown() && !data.IsType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"isType", data.IsType.ValueString())
	}
	if (!data.MetricType.IsUnknown() && !data.MetricType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"metricStyle", data.MetricType.ValueString())
	}
	if (!data.Mtu.IsUnknown() && !data.Mtu.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if (!data.Net.IsUnknown() && !data.Net.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"net", data.Net.ValueString())
	}
	if (!data.PassiveDefault.IsUnknown() && !data.PassiveDefault.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"passiveDflt", data.PassiveDefault.ValueString())
	}

	return nxos.Body{body}
}

func (data *ISISVRF) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
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
	if !data.BandwidthReference.IsNull() || all {
		data.BandwidthReference = types.Int64Value(res.Get(data.getClassName() + ".attributes.bwRef").Int())
	} else {
		data.BandwidthReference = types.Int64Null()
	}
	if !data.BanwidthReferenceUnit.IsNull() || all {
		data.BanwidthReferenceUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.bwRefUnit").String())
	} else {
		data.BanwidthReferenceUnit = types.StringNull()
	}
	if !data.IsType.IsNull() || all {
		data.IsType = types.StringValue(res.Get(data.getClassName() + ".attributes.isType").String())
	} else {
		data.IsType = types.StringNull()
	}
	if !data.MetricType.IsNull() || all {
		data.MetricType = types.StringValue(res.Get(data.getClassName() + ".attributes.metricStyle").String())
	} else {
		data.MetricType = types.StringNull()
	}
	if !data.Mtu.IsNull() || all {
		data.Mtu = types.Int64Value(res.Get(data.getClassName() + ".attributes.mtu").Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if !data.Net.IsNull() || all {
		data.Net = types.StringValue(res.Get(data.getClassName() + ".attributes.net").String())
	} else {
		data.Net = types.StringNull()
	}
	if !data.PassiveDefault.IsNull() || all {
		data.PassiveDefault = types.StringValue(res.Get(data.getClassName() + ".attributes.passiveDflt").String())
	} else {
		data.PassiveDefault = types.StringNull()
	}
}

func (data ISISVRF) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *ISISVRF) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/isis/inst-[%s]/dom-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InstanceName = types.StringValue(matches[1])
	data.Name = types.StringValue(matches[2])
}
