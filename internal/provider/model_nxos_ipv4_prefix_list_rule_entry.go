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

type IPv4PrefixListRuleEntry struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	RuleName  types.String `tfsdk:"rule_name"`
	Order     types.Int64  `tfsdk:"order"`
	Action    types.String `tfsdk:"action"`
	Criteria  types.String `tfsdk:"criteria"`
	Prefix    types.String `tfsdk:"prefix"`
	FromRange types.Int64  `tfsdk:"from_range"`
	ToRange   types.Int64  `tfsdk:"to_range"`
}

func (data IPv4PrefixListRuleEntry) getDn() string {
	return fmt.Sprintf("sys/rpm/pfxlistv4-[%s]/ent-[%v]", data.RuleName.ValueString(), data.Order.ValueInt64())
}

func (data IPv4PrefixListRuleEntry) getClassName() string {
	return "rtpfxEntry"
}

func (data IPv4PrefixListRuleEntry) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Order.IsUnknown() && !data.Order.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"order", strconv.FormatInt(data.Order.ValueInt64(), 10))
	}
	if (!data.Action.IsUnknown() && !data.Action.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"action", data.Action.ValueString())
	}
	if (!data.Criteria.IsUnknown() && !data.Criteria.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"criteria", data.Criteria.ValueString())
	}
	if (!data.Prefix.IsUnknown() && !data.Prefix.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pfx", data.Prefix.ValueString())
	}
	if (!data.FromRange.IsUnknown() && !data.FromRange.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"fromPfxLen", strconv.FormatInt(data.FromRange.ValueInt64(), 10))
	}
	if (!data.ToRange.IsUnknown() && !data.ToRange.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"toPfxLen", strconv.FormatInt(data.ToRange.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *IPv4PrefixListRuleEntry) fromBody(res gjson.Result, all bool) {
	if !data.Order.IsNull() || all {
		data.Order = types.Int64Value(res.Get(data.getClassName() + ".attributes.order").Int())
	} else {
		data.Order = types.Int64Null()
	}
	if !data.Action.IsNull() || all {
		data.Action = types.StringValue(res.Get(data.getClassName() + ".attributes.action").String())
	} else {
		data.Action = types.StringNull()
	}
	if !data.Criteria.IsNull() || all {
		data.Criteria = types.StringValue(res.Get(data.getClassName() + ".attributes.criteria").String())
	} else {
		data.Criteria = types.StringNull()
	}
	if !data.Prefix.IsNull() || all {
		data.Prefix = types.StringValue(res.Get(data.getClassName() + ".attributes.pfx").String())
	} else {
		data.Prefix = types.StringNull()
	}
	if !data.FromRange.IsNull() || all {
		data.FromRange = types.Int64Value(res.Get(data.getClassName() + ".attributes.fromPfxLen").Int())
	} else {
		data.FromRange = types.Int64Null()
	}
	if !data.ToRange.IsNull() || all {
		data.ToRange = types.Int64Value(res.Get(data.getClassName() + ".attributes.toPfxLen").Int())
	} else {
		data.ToRange = types.Int64Null()
	}
}

func (data IPv4PrefixListRuleEntry) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
