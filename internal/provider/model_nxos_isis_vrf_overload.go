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

type ISISVRFOverload struct {
	Device       types.String `tfsdk:"device"`
	Dn           types.String `tfsdk:"id"`
	InstanceName types.String `tfsdk:"instance_name"`
	Vrf          types.String `tfsdk:"vrf"`
	StartupTime  types.Int64  `tfsdk:"startup_time"`
}

func (data ISISVRFOverload) getDn() string {
	return fmt.Sprintf("sys/isis/inst-[%s]/dom-[%s]/overload", data.InstanceName.ValueString(), data.Vrf.ValueString())
}

func (data ISISVRFOverload) getClassName() string {
	return "isisOverload"
}

func (data ISISVRFOverload) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.StartupTime.IsUnknown() && !data.StartupTime.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"startupTime", strconv.FormatInt(data.StartupTime.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *ISISVRFOverload) fromBody(res gjson.Result, all bool) {
	if !data.StartupTime.IsNull() || all {
		data.StartupTime = types.Int64Value(res.Get(data.getClassName() + ".attributes.startupTime").Int())
	} else {
		data.StartupTime = types.Int64Null()
	}
}

func (data ISISVRFOverload) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
