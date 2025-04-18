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

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type OSPFv3Interface struct {
	Device               types.String `tfsdk:"device"`
	Dn                   types.String `tfsdk:"id"`
	InterfaceId          types.String `tfsdk:"interface_id"`
	AdvertiseSecondaries types.Bool   `tfsdk:"advertise_secondaries"`
	Area                 types.String `tfsdk:"area"`
	Bfd                  types.String `tfsdk:"bfd"`
	Cost                 types.Int64  `tfsdk:"cost"`
	DeadInterval         types.Int64  `tfsdk:"dead_interval"`
	HelloInterval        types.Int64  `tfsdk:"hello_interval"`
	NetworkType          types.String `tfsdk:"network_type"`
	Passive              types.String `tfsdk:"passive"`
	Priority             types.Int64  `tfsdk:"priority"`
}

func (data OSPFv3Interface) getDn() string {
	return fmt.Sprintf("sys/ospfv3/if-[%s]", data.InterfaceId.ValueString())
}

func (data OSPFv3Interface) getClassName() string {
	return "ospfv3If"
}

func (data OSPFv3Interface) toBody(statusReplace bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if statusReplace {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.InterfaceId.IsUnknown() && !data.InterfaceId.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"id", data.InterfaceId.ValueString())
	}
	if (!data.AdvertiseSecondaries.IsUnknown() && !data.AdvertiseSecondaries.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"advSecondary", strconv.FormatBool(data.AdvertiseSecondaries.ValueBool()))
	}
	if (!data.Area.IsUnknown() && !data.Area.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"area", data.Area.ValueString())
	}
	if (!data.Bfd.IsUnknown() && !data.Bfd.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"bfdCtrl", data.Bfd.ValueString())
	}
	if (!data.Cost.IsUnknown() && !data.Cost.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"cost", strconv.FormatInt(data.Cost.ValueInt64(), 10))
	}
	if (!data.DeadInterval.IsUnknown() && !data.DeadInterval.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"deadIntvl", strconv.FormatInt(data.DeadInterval.ValueInt64(), 10))
	}
	if (!data.HelloInterval.IsUnknown() && !data.HelloInterval.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"helloIntvl", strconv.FormatInt(data.HelloInterval.ValueInt64(), 10))
	}
	if (!data.NetworkType.IsUnknown() && !data.NetworkType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"nwT", data.NetworkType.ValueString())
	}
	if (!data.Passive.IsUnknown() && !data.Passive.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"passive", data.Passive.ValueString())
	}
	if (!data.Priority.IsUnknown() && !data.Priority.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"prio", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	}

	return nxos.Body{body}
}

func (data *OSPFv3Interface) fromBody(res gjson.Result, all bool) {
	if !data.InterfaceId.IsNull() || all {
		data.InterfaceId = types.StringValue(res.Get(data.getClassName() + ".attributes.id").String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if !data.AdvertiseSecondaries.IsNull() || all {
		data.AdvertiseSecondaries = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.advSecondary").String()))
	} else {
		data.AdvertiseSecondaries = types.BoolNull()
	}
	if !data.Area.IsNull() || all {
		data.Area = types.StringValue(res.Get(data.getClassName() + ".attributes.area").String())
	} else {
		data.Area = types.StringNull()
	}
	if !data.Bfd.IsNull() || all {
		data.Bfd = types.StringValue(res.Get(data.getClassName() + ".attributes.bfdCtrl").String())
	} else {
		data.Bfd = types.StringNull()
	}
	if !data.Cost.IsNull() || all {
		data.Cost = types.Int64Value(res.Get(data.getClassName() + ".attributes.cost").Int())
	} else {
		data.Cost = types.Int64Null()
	}
	if !data.DeadInterval.IsNull() || all {
		data.DeadInterval = types.Int64Value(res.Get(data.getClassName() + ".attributes.deadIntvl").Int())
	} else {
		data.DeadInterval = types.Int64Null()
	}
	if !data.HelloInterval.IsNull() || all {
		data.HelloInterval = types.Int64Value(res.Get(data.getClassName() + ".attributes.helloIntvl").Int())
	} else {
		data.HelloInterval = types.Int64Null()
	}
	if !data.NetworkType.IsNull() || all {
		data.NetworkType = types.StringValue(res.Get(data.getClassName() + ".attributes.nwT").String())
	} else {
		data.NetworkType = types.StringNull()
	}
	if !data.Passive.IsNull() || all {
		data.Passive = types.StringValue(res.Get(data.getClassName() + ".attributes.passive").String())
	} else {
		data.Passive = types.StringNull()
	}
	if !data.Priority.IsNull() || all {
		data.Priority = types.Int64Value(res.Get(data.getClassName() + ".attributes.prio").Int())
	} else {
		data.Priority = types.Int64Null()
	}
}

func (data OSPFv3Interface) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}

func (data *OSPFv3Interface) getIdsFromDn() {
	reString := strings.ReplaceAll("sys/ospfv3/if-[%s]", "%[1]s", ".+")
	reString = strings.ReplaceAll(reString, "%s", "(.+)")
	reString = strings.ReplaceAll(reString, "%v", "(.+)")
	reString = strings.ReplaceAll(reString, "[", "\\[")
	reString = strings.ReplaceAll(reString, "]", "\\]")
	re := regexp.MustCompile(reString)
	matches := re.FindStringSubmatch(data.Dn.ValueString())
	data.InterfaceId = types.StringValue(matches[1])
}
