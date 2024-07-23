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

type PIMStaticRP struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	VrfName types.String `tfsdk:"vrf_name"`
	Address types.String `tfsdk:"address"`
}

func (data PIMStaticRP) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/staticrp/rp-[%s]", data.VrfName.ValueString(), data.Address.ValueString())
}

func (data PIMStaticRP) getClassName() string {
	return "pimStaticRP"
}

func (data PIMStaticRP) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Address.IsUnknown() && !data.Address.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"addr", data.Address.ValueString())
	}

	return nxos.Body{body}
}

func (data *PIMStaticRP) fromBody(res gjson.Result, all bool) {
	if !data.Address.IsNull() || all {
		data.Address = types.StringValue(res.Get(data.getClassName() + ".attributes.addr").String())
	} else {
		data.Address = types.StringNull()
	}
}

func (data PIMStaticRP) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
