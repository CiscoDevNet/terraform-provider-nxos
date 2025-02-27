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

type RouteMapRuleEntrySetRegularCommunityItem struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	RuleName  types.String `tfsdk:"rule_name"`
	Order     types.Int64  `tfsdk:"order"`
	Community types.String `tfsdk:"community"`
}

func (data RouteMapRuleEntrySetRegularCommunityItem) getDn() string {
	return fmt.Sprintf("sys/rpm/rtmap-[%s]/ent-[%v]/sregcomm/item-[%s]", data.RuleName.ValueString(), data.Order.ValueInt64(), data.Community.ValueString())
}

func (data RouteMapRuleEntrySetRegularCommunityItem) getClassName() string {
	return "rtregcomItem"
}

func (data RouteMapRuleEntrySetRegularCommunityItem) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Community.IsUnknown() && !data.Community.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"community", data.Community.ValueString())
	}

	return nxos.Body{body}
}

func (data *RouteMapRuleEntrySetRegularCommunityItem) fromBody(res gjson.Result, all bool) {
	if !data.Community.IsNull() || all {
		data.Community = types.StringValue(res.Get(data.getClassName() + ".attributes.community").String())
	} else {
		data.Community = types.StringNull()
	}
}

func (data RouteMapRuleEntrySetRegularCommunityItem) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *RouteMapRuleEntrySetRegularCommunityItem) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/rpm/rtmap-[%s]/ent-[%v]/sregcomm/item-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.RuleName = types.StringValue(matches[1])
	data.Order = types.Int64Value(helpers.Must(strconv.ParseInt(matches[2], 10, 0)))
	data.Community = types.StringValue(matches[3])
}
