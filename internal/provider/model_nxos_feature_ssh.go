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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FeatureSSH struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	AdminState types.String `tfsdk:"admin_state"`
}

func (data FeatureSSH) getDn() string {
	return "sys/fm/ssh"
}

func (data FeatureSSH) getClassName() string {
	return "fmSsh"
}

func (data FeatureSSH) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.AdminState.IsUnknown() && !data.AdminState.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", data.AdminState.ValueString())
	}

	return nxos.Body{body}
}

func (data *FeatureSSH) fromBody(res gjson.Result, all bool) {
	if !data.AdminState.IsNull() || all {
		data.AdminState = types.StringValue(res.Get(data.getClassName() + ".attributes.adminSt").String())
	} else {
		data.AdminState = types.StringNull()
	}
}

func (data FeatureSSH) toDeleteBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes."+"adminSt", "enabled")

	return nxos.Body{body}
}
