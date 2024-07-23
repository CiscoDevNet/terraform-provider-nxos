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

type BGPAddressFamily struct {
	Device                    types.String `tfsdk:"device"`
	Dn                        types.String `tfsdk:"id"`
	Asn                       types.String `tfsdk:"asn"`
	Vrf                       types.String `tfsdk:"vrf"`
	AddressFamily             types.String `tfsdk:"address_family"`
	CriticalNexthopTimeout    types.Int64  `tfsdk:"critical_nexthop_timeout"`
	NonCriticalNexthopTimeout types.Int64  `tfsdk:"non_critical_nexthop_timeout"`
}

func (data BGPAddressFamily) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/af-[%s]", data.Vrf.ValueString(), data.AddressFamily.ValueString())
}

func (data BGPAddressFamily) getClassName() string {
	return "bgpDomAf"
}

func (data BGPAddressFamily) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.AddressFamily.IsUnknown() && !data.AddressFamily.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"type", data.AddressFamily.ValueString())
	}
	if (!data.CriticalNexthopTimeout.IsUnknown() && !data.CriticalNexthopTimeout.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"critNhTimeout", strconv.FormatInt(data.CriticalNexthopTimeout.ValueInt64(), 10))
	}
	if (!data.NonCriticalNexthopTimeout.IsUnknown() && !data.NonCriticalNexthopTimeout.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nonCritNhTimeout", strconv.FormatInt(data.NonCriticalNexthopTimeout.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *BGPAddressFamily) fromBody(res gjson.Result, all bool) {
	if !data.AddressFamily.IsNull() || all {
		data.AddressFamily = types.StringValue(res.Get(data.getClassName() + ".attributes.type").String())
	} else {
		data.AddressFamily = types.StringNull()
	}
	if !data.CriticalNexthopTimeout.IsNull() || all {
		data.CriticalNexthopTimeout = types.Int64Value(res.Get(data.getClassName() + ".attributes.critNhTimeout").Int())
	} else {
		data.CriticalNexthopTimeout = types.Int64Null()
	}
	if !data.NonCriticalNexthopTimeout.IsNull() || all {
		data.NonCriticalNexthopTimeout = types.Int64Value(res.Get(data.getClassName() + ".attributes.nonCritNhTimeout").Int())
	} else {
		data.NonCriticalNexthopTimeout = types.Int64Null()
	}
}

func (data BGPAddressFamily) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
