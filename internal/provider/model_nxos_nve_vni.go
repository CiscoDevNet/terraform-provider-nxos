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

type NVEVNI struct {
	Device                      types.String `tfsdk:"device"`
	Dn                          types.String `tfsdk:"id"`
	Vni                         types.Int64  `tfsdk:"vni"`
	AssociateVrf                types.Bool   `tfsdk:"associate_vrf"`
	MulticastGroup              types.String `tfsdk:"multicast_group"`
	MultisiteIngressReplication types.String `tfsdk:"multisite_ingress_replication"`
	SuppressArp                 types.String `tfsdk:"suppress_arp"`
}

func (data NVEVNI) getDn() string {
	return fmt.Sprintf("sys/eps/epId-[1]/nws/vni-[%v]", data.Vni.ValueInt64())
}

func (data NVEVNI) getClassName() string {
	return "nvoNw"
}

func (data NVEVNI) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Vni.IsUnknown() && !data.Vni.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"vni", strconv.FormatInt(data.Vni.ValueInt64(), 10))
	}
	if (!data.AssociateVrf.IsUnknown() && !data.AssociateVrf.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"associateVrfFlag", strconv.FormatBool(data.AssociateVrf.ValueBool()))
	}
	if (!data.MulticastGroup.IsUnknown() && !data.MulticastGroup.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"mcastGroup", data.MulticastGroup.ValueString())
	}
	if (!data.MultisiteIngressReplication.IsUnknown() && !data.MultisiteIngressReplication.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"multisiteIngRepl", data.MultisiteIngressReplication.ValueString())
	}
	if (!data.SuppressArp.IsUnknown() && !data.SuppressArp.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"suppressARP", data.SuppressArp.ValueString())
	}

	return nxos.Body{body}
}

func (data *NVEVNI) fromBody(res gjson.Result, all bool) {
	if !data.Vni.IsNull() || all {
		data.Vni = types.Int64Value(res.Get(data.getClassName() + ".attributes.vni").Int())
	} else {
		data.Vni = types.Int64Null()
	}
	if !data.AssociateVrf.IsNull() || all {
		data.AssociateVrf = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.associateVrfFlag").String()))
	} else {
		data.AssociateVrf = types.BoolNull()
	}
	if !data.MulticastGroup.IsNull() || all {
		data.MulticastGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.mcastGroup").String())
	} else {
		data.MulticastGroup = types.StringNull()
	}
	if !data.MultisiteIngressReplication.IsNull() || all {
		data.MultisiteIngressReplication = types.StringValue(res.Get(data.getClassName() + ".attributes.multisiteIngRepl").String())
	} else {
		data.MultisiteIngressReplication = types.StringNull()
	}
	if !data.SuppressArp.IsNull() || all {
		data.SuppressArp = types.StringValue(res.Get(data.getClassName() + ".attributes.suppressARP").String())
	} else {
		data.SuppressArp = types.StringNull()
	}
}

func (data NVEVNI) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
