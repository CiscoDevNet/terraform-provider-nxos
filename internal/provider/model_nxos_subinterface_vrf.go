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

type SubinterfaceVRF struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	InterfaceId types.String `tfsdk:"interface_id"`
	VrfDn       types.String `tfsdk:"vrf_dn"`
}

func (data SubinterfaceVRF) getDn() string {
	return fmt.Sprintf("sys/intf/encrtd-[%s]/rtvrfMbr", data.InterfaceId.ValueString())
}

func (data SubinterfaceVRF) getClassName() string {
	return "nwRtVrfMbr"
}

func (data SubinterfaceVRF) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.VrfDn.IsUnknown() && !data.VrfDn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"tDn", data.VrfDn.ValueString())
	}

	return nxos.Body{body}
}

func (data *SubinterfaceVRF) fromBody(res gjson.Result, all bool) {
	if !data.VrfDn.IsNull() || all {
		data.VrfDn = types.StringValue(res.Get(data.getClassName() + ".attributes.tDn").String())
	} else {
		data.VrfDn = types.StringNull()
	}
}

func (data SubinterfaceVRF) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *SubinterfaceVRF) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/intf/encrtd-[%s]/rtvrfMbr", "%[1]s", "")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InterfaceId = types.StringValue(matches[1])
}
