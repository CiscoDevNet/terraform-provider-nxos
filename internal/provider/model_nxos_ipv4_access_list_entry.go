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

type IPv4AccessListEntry struct {
	Device                  types.String `tfsdk:"device"`
	Dn                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	SequenceNumber          types.Int64  `tfsdk:"sequence_number"`
	Ack                     types.Bool   `tfsdk:"ack"`
	Action                  types.String `tfsdk:"action"`
	Dscp                    types.Int64  `tfsdk:"dscp"`
	DestinationAddressGroup types.String `tfsdk:"destination_address_group"`
	DestinationPort1        types.String `tfsdk:"destination_port_1"`
	DestinationPort2        types.String `tfsdk:"destination_port_2"`
	DestinationPortGroup    types.String `tfsdk:"destination_port_group"`
	DestinationPortMask     types.String `tfsdk:"destination_port_mask"`
	DestinationPortOperator types.String `tfsdk:"destination_port_operator"`
	DestinationPrefix       types.String `tfsdk:"destination_prefix"`
	DestinationPrefixLength types.String `tfsdk:"destination_prefix_length"`
	DestinationPrefixMask   types.String `tfsdk:"destination_prefix_mask"`
	Est                     types.Bool   `tfsdk:"est"`
	Fin                     types.Bool   `tfsdk:"fin"`
	Fragment                types.Bool   `tfsdk:"fragment"`
	HttpOptionType          types.String `tfsdk:"http_option_type"`
	IcmpCode                types.Int64  `tfsdk:"icmp_code"`
	IcmpType                types.Int64  `tfsdk:"icmp_type"`
	Logging                 types.Bool   `tfsdk:"logging"`
	PacketLength1           types.String `tfsdk:"packet_length_1"`
	PacketLength2           types.String `tfsdk:"packet_length_2"`
	PacketLengthOperator    types.String `tfsdk:"packet_length_operator"`
	Precedence              types.String `tfsdk:"precedence"`
	Protocol                types.String `tfsdk:"protocol"`
	ProtocolMask            types.String `tfsdk:"protocol_mask"`
	Psh                     types.Bool   `tfsdk:"psh"`
	Redirect                types.String `tfsdk:"redirect"`
	Remark                  types.String `tfsdk:"remark"`
	Rev                     types.Bool   `tfsdk:"rev"`
	Rst                     types.Bool   `tfsdk:"rst"`
	SourceAddressGroup      types.String `tfsdk:"source_address_group"`
	SourcePort1             types.String `tfsdk:"source_port_1"`
	SourcePort2             types.String `tfsdk:"source_port_2"`
	SourcePortGroup         types.String `tfsdk:"source_port_group"`
	SourcePortMask          types.String `tfsdk:"source_port_mask"`
	SourcePortOperator      types.String `tfsdk:"source_port_operator"`
	SourcePrefix            types.String `tfsdk:"source_prefix"`
	SourcePrefixLength      types.String `tfsdk:"source_prefix_length"`
	SourcePrefixMask        types.String `tfsdk:"source_prefix_mask"`
	Syn                     types.Bool   `tfsdk:"syn"`
	TimeRange               types.String `tfsdk:"time_range"`
	Ttl                     types.Int64  `tfsdk:"ttl"`
	Urg                     types.Bool   `tfsdk:"urg"`
	Vlan                    types.Int64  `tfsdk:"vlan"`
	Vni                     types.String `tfsdk:"vni"`
}

func (data IPv4AccessListEntry) getDn() string {
	return fmt.Sprintf("sys/acl/ipv4/name-[%s]/seq-[%v]", data.Name.ValueString(), data.SequenceNumber.ValueInt64())
}

func (data IPv4AccessListEntry) getClassName() string {
	return "ipv4aclACE"
}

func (data IPv4AccessListEntry) toBody(update bool) nxos.Body {
	body := ""
	body, _ = sjson.Set(body, data.getClassName()+".attributes", map[string]interface{}{})
	if update {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"status", "replaced")
	}
	if (!data.SequenceNumber.IsUnknown() && !data.SequenceNumber.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"seqNum", strconv.FormatInt(data.SequenceNumber.ValueInt64(), 10))
	}
	if (!data.Ack.IsUnknown() && !data.Ack.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ack", strconv.FormatBool(data.Ack.ValueBool()))
	}
	if (!data.Action.IsUnknown() && !data.Action.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"action", data.Action.ValueString())
	}
	if (!data.Dscp.IsUnknown() && !data.Dscp.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dscp", strconv.FormatInt(data.Dscp.ValueInt64(), 10))
	}
	if (!data.DestinationAddressGroup.IsUnknown() && !data.DestinationAddressGroup.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstAddrGroup", data.DestinationAddressGroup.ValueString())
	}
	if (!data.DestinationPort1.IsUnknown() && !data.DestinationPort1.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPort1", data.DestinationPort1.ValueString())
	}
	if (!data.DestinationPort2.IsUnknown() && !data.DestinationPort2.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPort2", data.DestinationPort2.ValueString())
	}
	if (!data.DestinationPortGroup.IsUnknown() && !data.DestinationPortGroup.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPortGroup", data.DestinationPortGroup.ValueString())
	}
	if (!data.DestinationPortMask.IsUnknown() && !data.DestinationPortMask.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPortMask", data.DestinationPortMask.ValueString())
	}
	if (!data.DestinationPortOperator.IsUnknown() && !data.DestinationPortOperator.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPortOp", data.DestinationPortOperator.ValueString())
	}
	if (!data.DestinationPrefix.IsUnknown() && !data.DestinationPrefix.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPrefix", data.DestinationPrefix.ValueString())
	}
	if (!data.DestinationPrefixLength.IsUnknown() && !data.DestinationPrefixLength.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPrefixLength", data.DestinationPrefixLength.ValueString())
	}
	if (!data.DestinationPrefixMask.IsUnknown() && !data.DestinationPrefixMask.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"dstPrefixMask", data.DestinationPrefixMask.ValueString())
	}
	if (!data.Est.IsUnknown() && !data.Est.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"est", strconv.FormatBool(data.Est.ValueBool()))
	}
	if (!data.Fin.IsUnknown() && !data.Fin.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"fin", strconv.FormatBool(data.Fin.ValueBool()))
	}
	if (!data.Fragment.IsUnknown() && !data.Fragment.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"fragment", strconv.FormatBool(data.Fragment.ValueBool()))
	}
	if (!data.HttpOptionType.IsUnknown() && !data.HttpOptionType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"httpOption", data.HttpOptionType.ValueString())
	}
	if (!data.IcmpCode.IsUnknown() && !data.IcmpCode.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"icmpCode", strconv.FormatInt(data.IcmpCode.ValueInt64(), 10))
	}
	if (!data.IcmpType.IsUnknown() && !data.IcmpType.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"icmpType", strconv.FormatInt(data.IcmpType.ValueInt64(), 10))
	}
	if (!data.Logging.IsUnknown() && !data.Logging.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"logging", strconv.FormatBool(data.Logging.ValueBool()))
	}
	if (!data.PacketLength1.IsUnknown() && !data.PacketLength1.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pktLen1", data.PacketLength1.ValueString())
	}
	if (!data.PacketLength2.IsUnknown() && !data.PacketLength2.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pktLen2", data.PacketLength2.ValueString())
	}
	if (!data.PacketLengthOperator.IsUnknown() && !data.PacketLengthOperator.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"pktLenOp", data.PacketLengthOperator.ValueString())
	}
	if (!data.Precedence.IsUnknown() && !data.Precedence.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"precedence", data.Precedence.ValueString())
	}
	if (!data.Protocol.IsUnknown() && !data.Protocol.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"protocol", data.Protocol.ValueString())
	}
	if (!data.ProtocolMask.IsUnknown() && !data.ProtocolMask.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"protocolMask", data.ProtocolMask.ValueString())
	}
	if (!data.Psh.IsUnknown() && !data.Psh.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"psh", strconv.FormatBool(data.Psh.ValueBool()))
	}
	if (!data.Redirect.IsUnknown() && !data.Redirect.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"redirect", data.Redirect.ValueString())
	}
	if (!data.Remark.IsUnknown() && !data.Remark.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"remark", data.Remark.ValueString())
	}
	if (!data.Rev.IsUnknown() && !data.Rev.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rev", strconv.FormatBool(data.Rev.ValueBool()))
	}
	if (!data.Rst.IsUnknown() && !data.Rst.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"rst", strconv.FormatBool(data.Rst.ValueBool()))
	}
	if (!data.SourceAddressGroup.IsUnknown() && !data.SourceAddressGroup.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcAddrGroup", data.SourceAddressGroup.ValueString())
	}
	if (!data.SourcePort1.IsUnknown() && !data.SourcePort1.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPort1", data.SourcePort1.ValueString())
	}
	if (!data.SourcePort2.IsUnknown() && !data.SourcePort2.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPort2", data.SourcePort2.ValueString())
	}
	if (!data.SourcePortGroup.IsUnknown() && !data.SourcePortGroup.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPortGroup", data.SourcePortGroup.ValueString())
	}
	if (!data.SourcePortMask.IsUnknown() && !data.SourcePortMask.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPortMask", data.SourcePortMask.ValueString())
	}
	if (!data.SourcePortOperator.IsUnknown() && !data.SourcePortOperator.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPortOp", data.SourcePortOperator.ValueString())
	}
	if (!data.SourcePrefix.IsUnknown() && !data.SourcePrefix.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPrefix", data.SourcePrefix.ValueString())
	}
	if (!data.SourcePrefixLength.IsUnknown() && !data.SourcePrefixLength.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPrefixLength", data.SourcePrefixLength.ValueString())
	}
	if (!data.SourcePrefixMask.IsUnknown() && !data.SourcePrefixMask.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"srcPrefixMask", data.SourcePrefixMask.ValueString())
	}
	if (!data.Syn.IsUnknown() && !data.Syn.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"syn", strconv.FormatBool(data.Syn.ValueBool()))
	}
	if (!data.TimeRange.IsUnknown() && !data.TimeRange.IsNull()) || false {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"timeRange", data.TimeRange.ValueString())
	}
	if (!data.Ttl.IsUnknown() && !data.Ttl.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"ttl", strconv.FormatInt(data.Ttl.ValueInt64(), 10))
	}
	if (!data.Urg.IsUnknown() && !data.Urg.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"urg", strconv.FormatBool(data.Urg.ValueBool()))
	}
	if (!data.Vlan.IsUnknown() && !data.Vlan.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"vlan", strconv.FormatInt(data.Vlan.ValueInt64(), 10))
	}
	if (!data.Vni.IsUnknown() && !data.Vni.IsNull()) || true {
		body, _ = sjson.Set(body, data.getClassName()+".attributes."+"vni", data.Vni.ValueString())
	}

	return nxos.Body{body}
}

func (data *IPv4AccessListEntry) fromBody(res gjson.Result, all bool) {
	if !data.SequenceNumber.IsNull() || all {
		data.SequenceNumber = types.Int64Value(res.Get(data.getClassName() + ".attributes.seqNum").Int())
	} else {
		data.SequenceNumber = types.Int64Null()
	}
	if !data.Ack.IsNull() || all {
		data.Ack = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.ack").String()))
	} else {
		data.Ack = types.BoolNull()
	}
	if !data.Action.IsNull() || all {
		data.Action = types.StringValue(res.Get(data.getClassName() + ".attributes.action").String())
	} else {
		data.Action = types.StringNull()
	}
	if !data.Dscp.IsNull() || all {
		data.Dscp = types.Int64Value(res.Get(data.getClassName() + ".attributes.dscp").Int())
	} else {
		data.Dscp = types.Int64Null()
	}
	if !data.DestinationAddressGroup.IsNull() || all {
		data.DestinationAddressGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.dstAddrGroup").String())
	} else {
		data.DestinationAddressGroup = types.StringNull()
	}
	if !data.DestinationPort1.IsNull() || all {
		data.DestinationPort1 = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPort1").String())
	} else {
		data.DestinationPort1 = types.StringNull()
	}
	if !data.DestinationPort2.IsNull() || all {
		data.DestinationPort2 = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPort2").String())
	} else {
		data.DestinationPort2 = types.StringNull()
	}
	if !data.DestinationPortGroup.IsNull() || all {
		data.DestinationPortGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPortGroup").String())
	} else {
		data.DestinationPortGroup = types.StringNull()
	}
	if !data.DestinationPortMask.IsNull() || all {
		data.DestinationPortMask = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPortMask").String())
	} else {
		data.DestinationPortMask = types.StringNull()
	}
	if !data.DestinationPortOperator.IsNull() || all {
		data.DestinationPortOperator = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPortOp").String())
	} else {
		data.DestinationPortOperator = types.StringNull()
	}
	if !data.DestinationPrefix.IsNull() || all {
		data.DestinationPrefix = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPrefix").String())
	} else {
		data.DestinationPrefix = types.StringNull()
	}
	if !data.DestinationPrefixLength.IsNull() || all {
		data.DestinationPrefixLength = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPrefixLength").String())
	} else {
		data.DestinationPrefixLength = types.StringNull()
	}
	if !data.DestinationPrefixMask.IsNull() || all {
		data.DestinationPrefixMask = types.StringValue(res.Get(data.getClassName() + ".attributes.dstPrefixMask").String())
	} else {
		data.DestinationPrefixMask = types.StringNull()
	}
	if !data.Est.IsNull() || all {
		data.Est = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.est").String()))
	} else {
		data.Est = types.BoolNull()
	}
	if !data.Fin.IsNull() || all {
		data.Fin = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.fin").String()))
	} else {
		data.Fin = types.BoolNull()
	}
	if !data.Fragment.IsNull() || all {
		data.Fragment = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.fragment").String()))
	} else {
		data.Fragment = types.BoolNull()
	}
	if !data.HttpOptionType.IsNull() || all {
		data.HttpOptionType = types.StringValue(res.Get(data.getClassName() + ".attributes.httpOption").String())
	} else {
		data.HttpOptionType = types.StringNull()
	}
	if !data.IcmpCode.IsNull() || all {
		data.IcmpCode = types.Int64Value(res.Get(data.getClassName() + ".attributes.icmpCode").Int())
	} else {
		data.IcmpCode = types.Int64Null()
	}
	if !data.IcmpType.IsNull() || all {
		data.IcmpType = types.Int64Value(res.Get(data.getClassName() + ".attributes.icmpType").Int())
	} else {
		data.IcmpType = types.Int64Null()
	}
	if !data.Logging.IsNull() || all {
		data.Logging = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.logging").String()))
	} else {
		data.Logging = types.BoolNull()
	}
	if !data.PacketLength1.IsNull() || all {
		data.PacketLength1 = types.StringValue(res.Get(data.getClassName() + ".attributes.pktLen1").String())
	} else {
		data.PacketLength1 = types.StringNull()
	}
	if !data.PacketLength2.IsNull() || all {
		data.PacketLength2 = types.StringValue(res.Get(data.getClassName() + ".attributes.pktLen2").String())
	} else {
		data.PacketLength2 = types.StringNull()
	}
	if !data.PacketLengthOperator.IsNull() || all {
		data.PacketLengthOperator = types.StringValue(res.Get(data.getClassName() + ".attributes.pktLenOp").String())
	} else {
		data.PacketLengthOperator = types.StringNull()
	}
	if !data.Precedence.IsNull() || all {
		data.Precedence = types.StringValue(res.Get(data.getClassName() + ".attributes.precedence").String())
	} else {
		data.Precedence = types.StringNull()
	}
	if !data.Protocol.IsNull() || all {
		data.Protocol = types.StringValue(res.Get(data.getClassName() + ".attributes.protocol").String())
	} else {
		data.Protocol = types.StringNull()
	}
	if !data.ProtocolMask.IsNull() || all {
		data.ProtocolMask = types.StringValue(res.Get(data.getClassName() + ".attributes.protocolMask").String())
	} else {
		data.ProtocolMask = types.StringNull()
	}
	if !data.Psh.IsNull() || all {
		data.Psh = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.psh").String()))
	} else {
		data.Psh = types.BoolNull()
	}
	if !data.Redirect.IsNull() || all {
		data.Redirect = types.StringValue(res.Get(data.getClassName() + ".attributes.redirect").String())
	} else {
		data.Redirect = types.StringNull()
	}
	if !data.Remark.IsNull() || all {
		data.Remark = types.StringValue(res.Get(data.getClassName() + ".attributes.remark").String())
	} else {
		data.Remark = types.StringNull()
	}
	if !data.Rev.IsNull() || all {
		data.Rev = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.rev").String()))
	} else {
		data.Rev = types.BoolNull()
	}
	if !data.Rst.IsNull() || all {
		data.Rst = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.rst").String()))
	} else {
		data.Rst = types.BoolNull()
	}
	if !data.SourceAddressGroup.IsNull() || all {
		data.SourceAddressGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.srcAddrGroup").String())
	} else {
		data.SourceAddressGroup = types.StringNull()
	}
	if !data.SourcePort1.IsNull() || all {
		data.SourcePort1 = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPort1").String())
	} else {
		data.SourcePort1 = types.StringNull()
	}
	if !data.SourcePort2.IsNull() || all {
		data.SourcePort2 = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPort2").String())
	} else {
		data.SourcePort2 = types.StringNull()
	}
	if !data.SourcePortGroup.IsNull() || all {
		data.SourcePortGroup = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPortGroup").String())
	} else {
		data.SourcePortGroup = types.StringNull()
	}
	if !data.SourcePortMask.IsNull() || all {
		data.SourcePortMask = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPortMask").String())
	} else {
		data.SourcePortMask = types.StringNull()
	}
	if !data.SourcePortOperator.IsNull() || all {
		data.SourcePortOperator = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPortOp").String())
	} else {
		data.SourcePortOperator = types.StringNull()
	}
	if !data.SourcePrefix.IsNull() || all {
		data.SourcePrefix = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPrefix").String())
	} else {
		data.SourcePrefix = types.StringNull()
	}
	if !data.SourcePrefixLength.IsNull() || all {
		data.SourcePrefixLength = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPrefixLength").String())
	} else {
		data.SourcePrefixLength = types.StringNull()
	}
	if !data.SourcePrefixMask.IsNull() || all {
		data.SourcePrefixMask = types.StringValue(res.Get(data.getClassName() + ".attributes.srcPrefixMask").String())
	} else {
		data.SourcePrefixMask = types.StringNull()
	}
	if !data.Syn.IsNull() || all {
		data.Syn = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.syn").String()))
	} else {
		data.Syn = types.BoolNull()
	}
	if !data.TimeRange.IsNull() || all {
		data.TimeRange = types.StringValue(res.Get(data.getClassName() + ".attributes.timeRange").String())
	} else {
		data.TimeRange = types.StringNull()
	}
	if !data.Ttl.IsNull() || all {
		data.Ttl = types.Int64Value(res.Get(data.getClassName() + ".attributes.ttl").Int())
	} else {
		data.Ttl = types.Int64Null()
	}
	if !data.Urg.IsNull() || all {
		data.Urg = types.BoolValue(helpers.ParseNxosBoolean(res.Get(data.getClassName() + ".attributes.urg").String()))
	} else {
		data.Urg = types.BoolNull()
	}
	if !data.Vlan.IsNull() || all {
		data.Vlan = types.Int64Value(res.Get(data.getClassName() + ".attributes.vlan").Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if !data.Vni.IsNull() || all {
		data.Vni = types.StringValue(res.Get(data.getClassName() + ".attributes.vni").String())
	} else {
		data.Vni = types.StringNull()
	}
}

func (data IPv4AccessListEntry) toDeleteBody() nxos.Body {
	body := ""

	return nxos.Body{body}
}
