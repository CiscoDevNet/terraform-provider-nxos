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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type IPv4AccessListPolicyIngressInterface struct {
	Device         types.String `tfsdk:"device"`
	Dn             types.String `tfsdk:"id"`
	InterfaceId    types.String `tfsdk:"interface_id"`
	AccessListName types.String `tfsdk:"access_list_name"`
}

func (data IPv4AccessListPolicyIngressInterface) getDn() string {
	return fmt.Sprintf("sys/acl/ipv4/policy/ingress/intf-[%s]", data.InterfaceId.ValueString())
}

func (data IPv4AccessListPolicyIngressInterface) getClassName() string {
	return "aclIf"
}

func (data IPv4AccessListPolicyIngressInterface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.InterfaceId.ValueString())
	}
	var attrs string
	attrs = ""
	if (!data.AccessListName.IsUnknown() && !data.AccessListName.IsNull()) || true {
		attrs, _ = sjson.Set(attrs, "name", data.AccessListName.ValueString())
	}
	body, _ = sjson.SetRaw(body, data.getClassName()+".children.-1.aclInst.attributes", attrs)

	return nxos.Body{body}
}

func (data *IPv4AccessListPolicyIngressInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	var r gjson.Result
	res.Get(data.getClassName() + ".children").ForEach(
		func(_, v gjson.Result) bool {
			key := v.Get("aclInst.attributes.rn").String()
			if key == "acl" {
				r = v
				return false
			}
			return true
		},
	)
	if !data.AccessListName.IsNull() || all {
		data.AccessListName = types.StringValue(r.Get("aclInst.attributes.name").String())
	} else {
		data.AccessListName = types.StringNull()
	}
}

func (data IPv4AccessListPolicyIngressInterface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *IPv4AccessListPolicyIngressInterface) getIdsFromDn() {
	reString := "sys/acl/ipv4/policy/ingress/intf-[%s]"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InterfaceId = types.StringValue(matches[1])
}
