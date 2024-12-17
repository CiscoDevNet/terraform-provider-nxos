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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DefaultQOSPolicyInterfaceInPolicyMap struct {
	Device        types.String `tfsdk:"device"`
	Dn            types.String `tfsdk:"id"`
	InterfaceId   types.String `tfsdk:"interface_id"`
	PolicyMapName types.String `tfsdk:"policy_map_name"`
}

func (data DefaultQOSPolicyInterfaceInPolicyMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/policy/in/intf-[%s]/pmap", data.InterfaceId.ValueString())
}

func (data DefaultQOSPolicyInterfaceInPolicyMap) getClassName() string {
	return "ipqosInst"
}

func (data DefaultQOSPolicyInterfaceInPolicyMap) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.PolicyMapName.IsUnknown() && !data.PolicyMapName.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.PolicyMapName.ValueString())
	}

	return nxos.Body{body}
}

func (data *DefaultQOSPolicyInterfaceInPolicyMap) fromBody(res gjson.Result, all bool) {
	if !data.PolicyMapName.IsNull() || all {
		data.PolicyMapName = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.PolicyMapName = types.StringNull()
	}
}

func (data DefaultQOSPolicyInterfaceInPolicyMap) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
