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

type BGPRouteRedistribution struct {
	Device           types.String `tfsdk:"device"`
	Dn               types.String `tfsdk:"id"`
	Asn              types.String `tfsdk:"asn"`
	Vrf              types.String `tfsdk:"vrf"`
	AddressFamily    types.String `tfsdk:"address_family"`
	Protocol         types.String `tfsdk:"protocol"`
	ProtocolInstance types.String `tfsdk:"protocol_instance"`
	RouteMap         types.String `tfsdk:"route_map"`
	Scope            types.String `tfsdk:"scope"`
	Srv6PrefixType   types.String `tfsdk:"srv6_prefix_type"`
}

func (data BGPRouteRedistribution) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/af-[%s]/interleak-[%s]-interleak-[%s]", data.Vrf.ValueString(), data.AddressFamily.ValueString(), data.Protocol.ValueString(), data.ProtocolInstance.ValueString())
}

func (data BGPRouteRedistribution) getClassName() string {
	return "bgpInterLeakP"
}

func (data BGPRouteRedistribution) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Protocol.IsUnknown() && !data.Protocol.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"proto", data.Protocol.ValueString())
	}
	if (!data.ProtocolInstance.IsUnknown() && !data.ProtocolInstance.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"inst", data.ProtocolInstance.ValueString())
	}
	if (!data.RouteMap.IsUnknown() && !data.RouteMap.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rtMap", data.RouteMap.ValueString())
	}
	if (!data.Scope.IsUnknown() && !data.Scope.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"scope", data.Scope.ValueString())
	}
	if (!data.Srv6PrefixType.IsUnknown() && !data.Srv6PrefixType.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srv6PrefixType", data.Srv6PrefixType.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPRouteRedistribution) fromBody(res gjson.Result, all bool) {
	if !data.Protocol.IsNull() || all {
		data.Protocol = types.StringValue(res.Get(data.getClassName() + ".attributes.proto").String())
	} else {
		data.Protocol = types.StringNull()
	}
	if !data.ProtocolInstance.IsNull() || all {
		data.ProtocolInstance = types.StringValue(res.Get(data.getClassName() + ".attributes.inst").String())
	} else {
		data.ProtocolInstance = types.StringNull()
	}
	if !data.RouteMap.IsNull() || all {
		data.RouteMap = types.StringValue(res.Get(data.getClassName() + ".attributes.rtMap").String())
	} else {
		data.RouteMap = types.StringNull()
	}
	if !data.Scope.IsNull() || all {
		data.Scope = types.StringValue(res.Get(data.getClassName() + ".attributes.scope").String())
	} else {
		data.Scope = types.StringNull()
	}
	if !data.Srv6PrefixType.IsNull() || all {
		data.Srv6PrefixType = types.StringValue(res.Get(data.getClassName() + ".attributes.srv6PrefixType").String())
	} else {
		data.Srv6PrefixType = types.StringNull()
	}
}

func (data BGPRouteRedistribution) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
