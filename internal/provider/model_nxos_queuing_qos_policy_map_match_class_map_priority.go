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
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type QueuingQOSPolicyMapMatchClassMapPriority struct {
	Device        types.String `tfsdk:"device"`
	Dn            types.String `tfsdk:"id"`
	PolicyMapName types.String `tfsdk:"policy_map_name"`
	ClassMapName  types.String `tfsdk:"class_map_name"`
	Level         types.Int64  `tfsdk:"level"`
}

func (data QueuingQOSPolicyMapMatchClassMapPriority) getDn() string {
	return fmt.Sprintf("sys/ipqos/queuing/p/name-[%s]/cmap-[%s]/prio", data.PolicyMapName.ValueString(), data.ClassMapName.ValueString())
}

func (data QueuingQOSPolicyMapMatchClassMapPriority) getClassName() string {
	return "ipqosPriority"
}

func (data QueuingQOSPolicyMapMatchClassMapPriority) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.Level.IsUnknown() && !data.Level.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"level", strconv.FormatInt(data.Level.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *QueuingQOSPolicyMapMatchClassMapPriority) fromBody(res gjson.Result, all bool) {
	if !data.Level.IsNull() || all {
		data.Level = types.Int64Value(res.Get(data.getClassName() + ".attributes.level").Int())
	} else {
		data.Level = types.Int64Null()
	}
}

func (data QueuingQOSPolicyMapMatchClassMapPriority) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *QueuingQOSPolicyMapMatchClassMapPriority) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ipqos/queuing/p/name-[%s]/cmap-[%s]/prio", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.PolicyMapName = types.StringValue(matches[1])
	data.ClassMapName = types.StringValue(matches[2])
}
