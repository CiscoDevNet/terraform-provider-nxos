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
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &IPv4AccessListEntryResource{}
var _ resource.ResourceWithImportState = &IPv4AccessListEntryResource{}

func NewIPv4AccessListEntryResource() resource.Resource {
	return &IPv4AccessListEntryResource{}
}

type IPv4AccessListEntryResource struct {
	clients map[string]*nxos.Client
}

func (r *IPv4AccessListEntryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_access_list_entry"
}

func (r *IPv4AccessListEntryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage IPv4 Access List Entries.", "ipv4aclACE", "Security%20and%20Policing/ipv4acl:ACE/").AddParents("ipv4_access_list").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access list name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"sequence_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Sequence number.").String,
				Required:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"ack": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP ACK flag.").String,
				Optional:            true,
			},
			"action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Action.").AddStringEnumDescription("invalid", "permit", "deny").AddDefaultValueDescription("invalid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("invalid"),
				Validators: []validator.String{
					stringvalidator.OneOf("invalid", "permit", "deny"),
				},
			},
			"dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match DSCP.").AddIntegerRangeDescription(0, 63).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"destination_address_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination address group.").String,
				Optional:            true,
			},
			"destination_port_1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("First destination port number or name.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"destination_port_2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Second destination port number or name.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"destination_port_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination port group.").String,
				Optional:            true,
			},
			"destination_port_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination port mask number or name.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"destination_port_operator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination port operator.").AddStringEnumDescription("none", "lt", "gt", "eq", "neq", "range").AddDefaultValueDescription("none").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
				Validators: []validator.String{
					stringvalidator.OneOf("none", "lt", "gt", "eq", "neq", "range"),
				},
			},
			"destination_prefix": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix.").String,
				Optional:            true,
			},
			"destination_prefix_length": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix length.").String,
				Optional:            true,
			},
			"destination_prefix_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination prefix mask.").String,
				Optional:            true,
			},
			"est": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP EST flag.").String,
				Optional:            true,
			},
			"fin": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP FIN flag.").String,
				Optional:            true,
			},
			"fragment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match non-initial fragment.").String,
				Optional:            true,
			},
			"http_option_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("HTTP option method.").AddStringEnumDescription("invalid", "get", "put", "head", "post", "delete", "trace", "connect").AddDefaultValueDescription("invalid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("invalid"),
				Validators: []validator.String{
					stringvalidator.OneOf("invalid", "get", "put", "head", "post", "delete", "trace", "connect"),
				},
			},
			"icmp_code": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ICMP code.").AddIntegerRangeDescription(0, 256).AddDefaultValueDescription("256").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(256),
				Validators: []validator.Int64{
					int64validator.Between(0, 256),
				},
			},
			"icmp_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("ICMP type.").AddIntegerRangeDescription(0, 256).AddDefaultValueDescription("256").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(256),
				Validators: []validator.Int64{
					int64validator.Between(0, 256),
				},
			},
			"logging": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log matches against ACL entry.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"packet_length_1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("First packet length. Either `invalid` or a number between 19 and 9210.").AddDefaultValueDescription("invalid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("invalid"),
			},
			"packet_length_2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Second packet length. Either `invalid` or a number between 19 and 9210.").AddDefaultValueDescription("invalid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("invalid"),
			},
			"packet_length_operator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Packet length operator.").AddStringEnumDescription("none", "lt", "gt", "eq", "neq", "range").AddDefaultValueDescription("none").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
				Validators: []validator.String{
					stringvalidator.OneOf("none", "lt", "gt", "eq", "neq", "range"),
				},
			},
			"precedence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Precedence. Either `unspecified` or a number between 0 and 7.").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol name or number.").AddStringEnumDescription("ip", "icmp", "igmp", "tcp", "udp", "gre", "esp", "ahp", "eigrp", "ospf", "nos", "pim", "pcp", "udf").String,
				Optional:            true,
			},
			"protocol_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol mask name or number.").AddStringEnumDescription("ip", "icmp", "igmp", "tcp", "udp", "gre", "esp", "ahp", "eigrp", "ospf", "nos", "pim", "pcp", "udf").String,
				Optional:            true,
			},
			"psh": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP PSH flag.").String,
				Optional:            true,
			},
			"redirect": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redirect action.").String,
				Optional:            true,
			},
			"remark": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACL comment.").String,
				Optional:            true,
			},
			"rev": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP REV flag.").String,
				Optional:            true,
			},
			"rst": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP RST flag.").String,
				Optional:            true,
			},
			"source_address_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source address group.").String,
				Optional:            true,
			},
			"source_port_1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("First source port name or number.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"source_port_2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Second source port name or number.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"source_port_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source port group.").String,
				Optional:            true,
			},
			"source_port_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source port mask name or number.").AddStringEnumDescription("echo", "discard", "daytime", "chargen", "ftp-data", "ftp", "telnet", "smtp", "time", "nameserver", "whois", "tacacs", "domain", "bootps", "bootpc", "tftp", "gopher", "finger", "www", "hostname", "pop2", "pop3", "sunrpc", "ident", "nntp", "ntp", "netbios-ns", "netbios-dgm", "netbios-ss", "snmp", "snmptrap", "xdmcp", "bgp", "irc", "dnsix", "mobile-ip", "pim-auto-rp", "isakmp", "biff", "exec", "who", "login", "syslog", "cmd", "lpd", "talk", "rip", "uucp", "klogin", "kshell", "drip", "non500-isakmp").String,
				Optional:            true,
			},
			"source_port_operator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source port operator.").AddStringEnumDescription("none", "lt", "gt", "eq", "neq", "range").AddDefaultValueDescription("none").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
				Validators: []validator.String{
					stringvalidator.OneOf("none", "lt", "gt", "eq", "neq", "range"),
				},
			},
			"source_prefix": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source prefix.").String,
				Optional:            true,
			},
			"source_prefix_length": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source prefix length.").String,
				Optional:            true,
			},
			"source_prefix_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source prefix mask.").String,
				Optional:            true,
			},
			"syn": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP SYN flag.").String,
				Optional:            true,
			},
			"time_range": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Time range name.").String,
				Optional:            true,
			},
			"ttl": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("TTL.").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"urg": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Match TCP URG flag.").String,
				Optional:            true,
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("VLAN ID.").AddIntegerRangeDescription(0, 4095).AddDefaultValueDescription("4095").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(4095),
				Validators: []validator.Int64{
					int64validator.Between(0, 4095),
				},
			},
			"vni": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NVE VNI ID. Either `invalid` or a number between 0 and 16777216.").AddDefaultValueDescription("invalid").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("invalid"),
			},
		},
	}
}

func (r *IPv4AccessListEntryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *IPv4AccessListEntryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPv4AccessListEntry

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody(false)
	_, err := r.clients[plan.Device.ValueString()].Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *IPv4AccessListEntryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPv4AccessListEntry

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
	res, err := r.clients[state.Device.ValueString()].GetDn(state.Dn.ValueString(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}
	state.fromBody(res, imp)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *IPv4AccessListEntryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPv4AccessListEntry

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody(true)

	_, err := r.clients[plan.Device.ValueString()].Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IPv4AccessListEntryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPv4AccessListEntry

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	body := state.toDeleteBody()

	if len(body.Str) > 0 {
		_, err := r.clients[state.Device.ValueString()].Post(state.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := r.clients[state.Device.ValueString()].DeleteDn(state.Dn.ValueString())
		if err != nil {
			errCode := res.Get("imdata.0.error.attributes.code").Str
			// Ignore errors of type "Cannot delete object"
			if errCode != "1" && errCode != "107" {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *IPv4AccessListEntryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
