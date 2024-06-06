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

type DefaultQOSPolicyMapMatchClassMapSetQOSGroup struct {
	Device        types.String `tfsdk:"device"`
	Dn            types.String `tfsdk:"id"`
	PolicyMapName types.String `tfsdk:"policy_map_name"`
	ClassMapName  types.String `tfsdk:"class_map_name"`
	QosGroupId    types.Int64  `tfsdk:"qos_group_id"`
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/p/name-[%s]/cmap-[%s]/setGrp", data.PolicyMapName.ValueString(), data.ClassMapName.ValueString())
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) getClassName() string {
	return "ipqosSetQoSGrp"
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.QosGroupId.IsUnknown() && !data.QosGroupId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", strconv.FormatInt(data.QosGroupId.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *DefaultQOSPolicyMapMatchClassMapSetQOSGroup) fromBody(res gjson.Result, all bool) {
	if !data.QosGroupId.IsNull() || all {
		data.QosGroupId = types.Int64Value(res.Get(data.getClassName() + ".attributes.id").Int())
	} else {
		data.QosGroupId = types.Int64Null()
	}
}

func (data DefaultQOSPolicyMapMatchClassMapSetQOSGroup) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
