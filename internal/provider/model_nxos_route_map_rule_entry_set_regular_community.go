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

type RouteMapRuleEntrySetRegularCommunity struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	RuleName    types.String `tfsdk:"rule_name"`
	Order       types.Int64  `tfsdk:"order"`
	Additive    types.String `tfsdk:"additive"`
	NoCommunity types.String `tfsdk:"no_community"`
	SetCriteria types.String `tfsdk:"set_criteria"`
}

func (data RouteMapRuleEntrySetRegularCommunity) getDn() string {
	return fmt.Sprintf("sys/rpm/rtmap-[%s]/ent-[%v]/sregcomm", data.RuleName.ValueString(), data.Order.ValueInt64())
}

func (data RouteMapRuleEntrySetRegularCommunity) getClassName() string {
	return "rtmapSetRegComm"
}

func (data RouteMapRuleEntrySetRegularCommunity) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Additive.IsUnknown() && !data.Additive.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"additive", data.Additive.ValueString())
	}
	if (!data.NoCommunity.IsUnknown() && !data.NoCommunity.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"noCommAttr", data.NoCommunity.ValueString())
	}
	if (!data.SetCriteria.IsUnknown() && !data.SetCriteria.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"setCriteria", data.SetCriteria.ValueString())
	}

	return nxos.Body{body}
}

func (data *RouteMapRuleEntrySetRegularCommunity) fromBody(res gjson.Result, all bool) {
	if !data.Additive.IsNull() || all {
		data.Additive = types.StringValue(res.Get(data.getClassName() + ".attributes.additive").String())
	} else {
		data.Additive = types.StringNull()
	}
	if !data.NoCommunity.IsNull() || all {
		data.NoCommunity = types.StringValue(res.Get(data.getClassName() + ".attributes.noCommAttr").String())
	} else {
		data.NoCommunity = types.StringNull()
	}
	if !data.SetCriteria.IsNull() || all {
		data.SetCriteria = types.StringValue(res.Get(data.getClassName() + ".attributes.setCriteria").String())
	} else {
		data.SetCriteria = types.StringNull()
	}
}

func (data RouteMapRuleEntrySetRegularCommunity) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *RouteMapRuleEntrySetRegularCommunity) getIdsFromDn() {
	reString := "sys/rpm/rtmap-[%s]/ent-[%v]/sregcomm"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.RuleName = types.StringValue(matches[1])
	data.Order = types.Int64Value(helpers.Must(strconv.ParseInt(matches[2], 10, 0)))
}
