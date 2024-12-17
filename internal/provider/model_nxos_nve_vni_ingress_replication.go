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

type NVEVNIIngressReplication struct {
	Device   types.String `tfsdk:"device"`
	Dn       types.String `tfsdk:"id"`
	Vni      types.Int64  `tfsdk:"vni"`
	Protocol types.String `tfsdk:"protocol"`
}

func (data NVEVNIIngressReplication) getDn() string {
	return fmt.Sprintf("sys/eps/epId-[1]/nws/vni-[%v]/IngRepl", data.Vni.ValueInt64())
}

func (data NVEVNIIngressReplication) getClassName() string {
	return "nvoIngRepl"
}

func (data NVEVNIIngressReplication) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Protocol.IsUnknown() && !data.Protocol.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"proto", data.Protocol.ValueString())
	}

	return nxos.Body{body}
}

func (data *NVEVNIIngressReplication) fromBody(res gjson.Result, all bool) {
	if !data.Protocol.IsNull() || all {
		data.Protocol = types.StringValue(res.Get(data.getClassName() + ".attributes.proto").String())
	} else {
		data.Protocol = types.StringNull()
	}
}

func (data NVEVNIIngressReplication) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
