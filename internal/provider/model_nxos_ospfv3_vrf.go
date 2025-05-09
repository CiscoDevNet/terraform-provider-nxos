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

type OSPFv3VRF struct {
	Device                 types.String `tfsdk:"device"`
	Dn                     types.String `tfsdk:"id"`
	InstanceName           types.String `tfsdk:"instance_name"`
	Name                   types.String `tfsdk:"name"`
	AdminState             types.String `tfsdk:"admin_state"`
	BandwidthReference     types.Int64  `tfsdk:"bandwidth_reference"`
	BandwidthReferenceUnit types.String `tfsdk:"bandwidth_reference_unit"`
	RouterId               types.String `tfsdk:"router_id"`
	BfdControl             types.Bool   `tfsdk:"bfd_control"`
}

func (data OSPFv3VRF) getDn() string {
	return fmt.Sprintf("sys/ospfv3/inst-[%s]/dom-[%s]", data.InstanceName.ValueString(), data.Name.ValueString())
}

func (data OSPFv3VRF) getClassName() string {
	return "ospfv3Dom"
}

func (data OSPFv3VRF) toBody(statusReplace bool) nxos.Body {
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
	if (!data.BandwidthReference.IsUnknown() && !data.BandwidthReference.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bwRef", strconv.FormatInt(data.BandwidthReference.ValueInt64(), 10))
	}
	if (!data.BandwidthReferenceUnit.IsUnknown() && !data.BandwidthReferenceUnit.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bwRefUnit", data.BandwidthReferenceUnit.ValueString())
	}
	if (!data.RouterId.IsUnknown() && !data.RouterId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rtrId", data.RouterId.ValueString())
	}
	if (!data.BfdControl.IsUnknown() && !data.BfdControl.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bfdCtrl", strconv.FormatBool(data.BfdControl.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *OSPFv3VRF) fromBody(res gjson.Result, all bool) {
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
	if !data.BandwidthReference.IsNull() || all {
		data.BandwidthReference = types.Int64Value(res.Get(data.getClassName() + ".attributes.bwRef").Int())
	} else {
		data.BandwidthReference = types.Int64Null()
	}
	if !data.BandwidthReferenceUnit.IsNull() || all {
		data.BandwidthReferenceUnit = types.StringValue(res.Get(data.getClassName() + ".attributes.bwRefUnit").String())
	} else {
		data.BandwidthReferenceUnit = types.StringNull()
	}
	if !data.RouterId.IsNull() || all {
		data.RouterId = types.StringValue(res.Get(data.getClassName() + ".attributes.rtrId").String())
	} else {
		data.RouterId = types.StringNull()
	}
	if !data.BfdControl.IsNull() || all {
		data.BfdControl = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.bfdCtrl").String()))
	} else {
		data.BfdControl = types.BoolNull()
	}
}

func (data OSPFv3VRF) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *OSPFv3VRF) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ospfv3/inst-[%s]/dom-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InstanceName = types.StringValue(matches[1])
	data.Name = types.StringValue(matches[2])
}
