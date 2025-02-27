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

type OSPFArea struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	InstanceName       types.String `tfsdk:"instance_name"`
	VrfName            types.String `tfsdk:"vrf_name"`
	AreaId             types.String `tfsdk:"area_id"`
	AuthenticationType types.String `tfsdk:"authentication_type"`
	Cost               types.Int64  `tfsdk:"cost"`
	Type               types.String `tfsdk:"type"`
}

func (data OSPFArea) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/area-[%s]", data.InstanceName.ValueString(), data.VrfName.ValueString(), data.AreaId.ValueString())
}

func (data OSPFArea) getClassName() string {
	return "ospfArea"
}

func (data OSPFArea) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.AreaId.IsUnknown() && !data.AreaId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.AreaId.ValueString())
	}
	if (!data.AuthenticationType.IsUnknown() && !data.AuthenticationType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"authType", data.AuthenticationType.ValueString())
	}
	if (!data.Cost.IsUnknown() && !data.Cost.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cost", strconv.FormatInt(data.Cost.ValueInt64(), 10))
	}
	if (!data.Type.IsUnknown() && !data.Type.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.Type.ValueString())
	}

	return nxos.Body{body}
}

func (data *OSPFArea) fromBody(res gjson.Result, all bool) {
	if !data.AreaId.IsNull() || all {
		data.AreaId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.AreaId = types.StringNull()
	}
	if !data.AuthenticationType.IsNull() || all {
		data.AuthenticationType = types.StringValue(res.Get(data.getClassName() + ".attributes.authType").String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if !data.Cost.IsNull() || all {
		data.Cost = types.Int64Value(res.Get(data.getClassName() + ".attributes.cost").Int())
	} else {
		data.Cost = types.Int64Null()
	}
	if !data.Type.IsNull() || all {
		data.Type = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.Type = types.StringNull()
	}
}

func (data OSPFArea) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *OSPFArea) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ospf/inst-[%s]/dom-[%s]/area-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InstanceName = types.StringValue(matches[1])
	data.VrfName = types.StringValue(matches[2])
	data.AreaId = types.StringValue(matches[3])
}
