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

type BGPRouteControl struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	Asn                types.String `tfsdk:"asn"`
	Vrf                types.String `tfsdk:"vrf"`
	EnforceFirstAs     types.String `tfsdk:"enforce_first_as"`
	FibAccelerate      types.String `tfsdk:"fib_accelerate"`
	LogNeighborChanges types.String `tfsdk:"log_neighbor_changes"`
	SuppressRoutes     types.String `tfsdk:"suppress_routes"`
}

func (data BGPRouteControl) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/rtctrl", data.Vrf.ValueString())
}

func (data BGPRouteControl) getClassName() string {
	return "bgpRtCtrl"
}

func (data BGPRouteControl) toBody() nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if (!data.EnforceFirstAs.IsUnknown() && !data.EnforceFirstAs.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"enforceFirstAs", data.EnforceFirstAs.ValueString())
	}
	if (!data.FibAccelerate.IsUnknown() && !data.FibAccelerate.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"fibAccelerate", data.FibAccelerate.ValueString())
	}
	if (!data.LogNeighborChanges.IsUnknown() && !data.LogNeighborChanges.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"logNeighborChanges", data.LogNeighborChanges.ValueString())
	}
	if (!data.SuppressRoutes.IsUnknown() && !data.SuppressRoutes.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"supprRt", data.SuppressRoutes.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPRouteControl) fromBody(res gjson.Result, all bool) {
	if !data.EnforceFirstAs.IsNull() || all {
		data.EnforceFirstAs = types.StringValue(res.Get(data.getClassName() + ".attributes.enforceFirstAs").String())
	} else {
		data.EnforceFirstAs = types.StringNull()
	}
	if !data.FibAccelerate.IsNull() || all {
		data.FibAccelerate = types.StringValue(res.Get(data.getClassName() + ".attributes.fibAccelerate").String())
	} else {
		data.FibAccelerate = types.StringNull()
	}
	if !data.LogNeighborChanges.IsNull() || all {
		data.LogNeighborChanges = types.StringValue(res.Get(data.getClassName() + ".attributes.logNeighborChanges").String())
	} else {
		data.LogNeighborChanges = types.StringNull()
	}
	if !data.SuppressRoutes.IsNull() || all {
		data.SuppressRoutes = types.StringValue(res.Get(data.getClassName() + ".attributes.supprRt").String())
	} else {
		data.SuppressRoutes = types.StringNull()
	}
}

func (data BGPRouteControl) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
