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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Ethernet struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	Mtu                types.Int64  `tfsdk:"mtu"`
	DefaultAdminStatus types.String `tfsdk:"default_admin_status"`
}

func (data Ethernet) getDn() string {
	return "sys/ethpm/inst"
}

func (data Ethernet) getClassName() string {
	return "ethpmInst"
}

func (data Ethernet) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Mtu.IsUnknown() && !data.Mtu.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"systemJumboMtu", strconv.FormatInt(data.Mtu.ValueInt64(), 10))
	}
	if (!data.DefaultAdminStatus.IsUnknown() && !data.DefaultAdminStatus.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"systemDefaultAdminSt", data.DefaultAdminStatus.ValueString())
	}

	return nxos.Body{body}
}

func (data *Ethernet) fromBody(res gjson.Result, all bool) {
	if !data.Mtu.IsNull() || all {
		data.Mtu = types.Int64Value(res.Get(data.getClassName() + ".attributes.systemJumboMtu").Int())
	} else {
		data.Mtu = types.Int64Null()
	}
	if !data.DefaultAdminStatus.IsNull() || all {
		data.DefaultAdminStatus = types.StringValue(res.Get(data.getClassName() + ".attributes.systemDefaultAdminSt").String())
	} else {
		data.DefaultAdminStatus = types.StringNull()
	}
}

func (data Ethernet) toDeleteBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"systemJumboMtu", strconv.FormatInt(9216, 10))
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"systemDefaultAdminSt", "DME_UNSET_PROPERTY_MARKER")

	return nxos.Body{body}
}
