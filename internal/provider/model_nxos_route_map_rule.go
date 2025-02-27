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

type RouteMapRule struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
}

func (data RouteMapRule) getDn() string {
	return fmt.Sprintf("sys/rpm/rtmap-[%s]", data.Name.ValueString())
}

func (data RouteMapRule) getClassName() string {
	return "rtmapRule"
}

func (data RouteMapRule) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Name.IsUnknown() && !data.Name.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"name", data.Name.ValueString())
	}

	return nxos.Body{body}
}

func (data *RouteMapRule) fromBody(res gjson.Result, all bool) {
	if !data.Name.IsNull() || all {
		data.Name = types.StringValue(res.Get(data.getClassName() + ".attributes.name").String())
	} else {
		data.Name = types.StringNull()
	}
}

func (data RouteMapRule) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *RouteMapRule) getIdsFromDn() {
	reString := "sys/rpm/rtmap-[%s]"
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.Name = types.StringValue(matches[1])
}
