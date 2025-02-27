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

type PortChannelInterfaceMember struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	InterfaceId types.String `tfsdk:"interface_id"`
	InterfaceDn types.String `tfsdk:"interface_dn"`
	Force       types.Bool   `tfsdk:"force"`
}

func (data PortChannelInterfaceMember) getDn() string {
	return fmt.Sprintf("sys/intf/aggr-[%s]/rsmbrIfs-[%s]", data.InterfaceId.ValueString(), data.InterfaceDn.ValueString())
}

func (data PortChannelInterfaceMember) getClassName() string {
	return "pcRsMbrIfs"
}

func (data PortChannelInterfaceMember) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceDn.IsUnknown() && !data.InterfaceDn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tDn", data.InterfaceDn.ValueString())
	}
	if (!data.Force.IsUnknown() && !data.Force.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"isMbrForce", strconv.FormatBool(data.Force.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *PortChannelInterfaceMember) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceDn.IsNull() || all {
		data.InterfaceDn = types.StringValue(res.Get(data.getClassName() + ".attributes.tDn").String())
	} else {
		data.InterfaceDn = types.StringNull()
	}
	if !data.Force.IsNull() || all {
		data.Force = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.isMbrForce").String()))
	} else {
		data.Force = types.BoolNull()
	}
}

func (data PortChannelInterfaceMember) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *PortChannelInterfaceMember) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/intf/aggr-[%s]/rsmbrIfs-[%s]", "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InterfaceId = types.StringValue(matches[1])
	data.InterfaceDn = types.StringValue(matches[2])
}
