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
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type BGPPeerLocalASN struct {
	Device         types.String `tfsdk:"device"`
	Dn             types.String `tfsdk:"id"`
	AsnPropagation types.String `tfsdk:"asn_propagation"`
	LocalAsn       types.String `tfsdk:"local_asn"`
	Vrf            types.String `tfsdk:"vrf"`
	Address        types.String `tfsdk:"address"`
}

func (data BGPPeerLocalASN) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]/localasn", data.Vrf.ValueString(), data.Address.ValueString())
}

func (data BGPPeerLocalASN) getClassName() string {
	return "bgpLocalAsn"
}

func (data BGPPeerLocalASN) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.AsnPropagation.IsUnknown() && !data.AsnPropagation.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"asnPropagate", data.AsnPropagation.ValueString())
	}
	if (!data.LocalAsn.IsUnknown() && !data.LocalAsn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"localAsn", data.LocalAsn.ValueString())
	}

	return nxos.Body{body}
}

func (data *BGPPeerLocalASN) fromBody(res gjson.Result, all bool) {
	if !data.AsnPropagation.IsNull() || all {
		data.AsnPropagation = types.StringValue(res.Get(data.getClassName() + ".attributes.asnPropagate").String())
	} else {
		data.AsnPropagation = types.StringNull()
	}
	if !data.LocalAsn.IsNull() || all {
		data.LocalAsn = types.StringValue(res.Get(data.getClassName() + ".attributes.localAsn").String())
	} else {
		data.LocalAsn = types.StringNull()
	}
}

func (data BGPPeerLocalASN) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *BGPPeerLocalASN) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/bgp/inst/dom-[%s]/peer-[%s]/localasn", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Vrf = types.StringValue(matches[1])
	data.Address = types.StringValue(matches[2])
}
