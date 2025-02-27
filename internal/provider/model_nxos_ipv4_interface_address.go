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

type IPv4InterfaceAddress struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	Vrf         types.String `tfsdk:"vrf"`
	InterfaceId types.String `tfsdk:"interface_id"`
	Address     types.String `tfsdk:"address"`
	Type        types.String `tfsdk:"type"`
	Tag         types.Int64  `tfsdk:"tag"`
}

func (data IPv4InterfaceAddress) getDn() string {
	return fmt.Sprintf("sys/ipv4/inst/dom-[%s]/if-[%s]/addr-[%s]", data.Vrf.ValueString(), data.InterfaceId.ValueString(), data.Address.ValueString())
}

func (data IPv4InterfaceAddress) getClassName() string {
	return "ipv4Addr"
}

func (data IPv4InterfaceAddress) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"addr", data.Address.ValueString())
	}
	if (!data.Type.IsUnknown() && !data.Type.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.Type.ValueString())
	}
	if (!data.Tag.IsUnknown() && !data.Tag.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tag", strconv.FormatInt(data.Tag.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *IPv4InterfaceAddress) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.addr").String())
	} else {
		data.Address = types.StringNull()
	}
	if !data.Type.IsNull() || all {
		data.Type = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.Type = types.StringNull()
	}
	if !data.Tag.IsNull() || all {
		data.Tag = types.Int64Value(res.Get(data.getClassName() + ".attributes.tag").Int())
	} else {
		data.Tag = types.Int64Null()
	}
}

func (data IPv4InterfaceAddress) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *IPv4InterfaceAddress) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ipv4/inst/dom-[%s]/if-[%s]/addr-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.InterfaceId = types.StringValue(matches[2])
	data.Address = types.StringValue(matches[3])
}
