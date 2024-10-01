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

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type PIMInterface struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	VrfName     types.String `tfsdk:"vrf_name"`
	InterfaceId types.String `tfsdk:"interface_id"`
	AdminState  types.String `tfsdk:"admin_state"`
	Bfd         types.String `tfsdk:"bfd"`
	DrPriority  types.Int64  `tfsdk:"dr_priority"`
	Passive     types.Bool   `tfsdk:"passive"`
	SparseMode  types.Bool   `tfsdk:"sparse_mode"`
}

func (data PIMInterface) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/if-[%s]", data.VrfName.ValueString(), data.InterfaceId.ValueString())
}

func (data PIMInterface) getClassName() string {
	return "pimIf"
}

func (data PIMInterface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}
	if (!data.Bfd.IsUnknown() && !data.Bfd.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bfdInst", data.Bfd.ValueString())
	}
	if (!data.DrPriority.IsUnknown() && !data.DrPriority.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"drPrio", strconv.FormatInt(data.DrPriority.ValueInt64(), 10))
	}
	if (!data.Passive.IsUnknown() && !data.Passive.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"passive", strconv.FormatBool(data.Passive.ValueBool()))
	}
	if (!data.SparseMode.IsUnknown() && !data.SparseMode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pimSparseMode", strconv.FormatBool(data.SparseMode.ValueBool()))
	}

	return nxos.Body{body}
}

func (data *PIMInterface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
	if !data.Bfd.IsNull() || all {
		data.Bfd = types.StringValue(res.Get(data.getClassName() + ".attributes.bfdInst").String())
	} else {
		data.Bfd = types.StringNull()
	}
	if !data.DrPriority.IsNull() || all {
		data.DrPriority = types.Int64Value(res.Get(data.getClassName() + ".attributes.drPrio").Int())
	} else {
		data.DrPriority = types.Int64Null()
	}
	if !data.Passive.IsNull() || all {
		data.Passive = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.passive").String()))
	} else {
		data.Passive = types.BoolNull()
	}
	if !data.SparseMode.IsNull() || all {
		data.SparseMode = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.pimSparseMode").String()))
	} else {
		data.SparseMode = types.BoolNull()
	}
}

func (data PIMInterface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
