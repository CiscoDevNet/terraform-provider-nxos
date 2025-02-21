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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PortChannelInterfaceResource{}
var _ resource.ResourceWithImportState = &PortChannelInterfaceResource{}

func NewPortChannelInterfaceResource() resource.Resource {
	return &PortChannelInterfaceResource{}
}

type PortChannelInterfaceResource struct {
	clients map[string]*nxos.Client
}

func (r *PortChannelInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_port_channel_interface"
}

func (r *PortChannelInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage a port-channel interface.", "pcAggrIf", "Interfaces/pc:AggrIf/").AddChildren("port_channel_interface_vrf", "port_channel_interface_member").String,

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
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `po1`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"port_channel_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The aggregated interface protocol channel mode.").AddStringEnumDescription("on", "static", "active", "passive", "mac-pin").AddDefaultValueDescription("on").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("on"),
				Validators: []validator.String{
					stringvalidator.OneOf("on", "static", "active", "passive", "mac-pin"),
				},
			},
			"minimum_links": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum links.").AddIntegerRangeDescription(1, 32).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.Between(1, 32),
				},
			},
			"maximum_links": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maximum links.").AddIntegerRangeDescription(1, 32).AddDefaultValueDescription("32").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(32),
				Validators: []validator.Int64{
					int64validator.Between(1, 32),
				},
			},
			"suspend_individual": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Suspend Individual Port.").AddStringEnumDescription("enable", "disable").AddDefaultValueDescription("enable").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("enable"),
				Validators: []validator.String{
					stringvalidator.OneOf("enable", "disable"),
				},
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("vlan-1"),
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port state.").AddStringEnumDescription("up", "down").AddDefaultValueDescription("up").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("up"),
				Validators: []validator.String{
					stringvalidator.OneOf("up", "down"),
				},
			},
			"auto_negotiation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port auto-negotiation.").AddStringEnumDescription("on", "off", "25G").AddDefaultValueDescription("on").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("on"),
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "25G"),
				},
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The bandwidth parameter for a routed interface, port channel, or subinterface.").AddIntegerRangeDescription(0, 3200000000).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 3200000000),
				},
			},
			"delay": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port delay time.").AddIntegerRangeDescription(1, 16777215).AddDefaultValueDescription("1").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.Between(1, 16777215),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface description.").String,
				Optional:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duplex.").AddStringEnumDescription("auto", "full", "half").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("auto"),
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "full", "half"),
				},
			},
			"layer": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port layer.").AddStringEnumDescription("Layer2", "Layer3").AddDefaultValueDescription("Layer2").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("Layer2"),
				Validators: []validator.String{
					stringvalidator.OneOf("Layer2", "Layer3"),
				},
			},
			"link_logging": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative link logging.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				Validators: []validator.String{
					stringvalidator.OneOf("default", "enable", "disable"),
				},
			},
			"medium": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port medium type.").AddStringEnumDescription("broadcast", "p2p").AddDefaultValueDescription("broadcast").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("broadcast"),
				Validators: []validator.String{
					stringvalidator.OneOf("broadcast", "p2p"),
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port mode.").AddStringEnumDescription("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag").AddDefaultValueDescription("access").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("access"),
				Validators: []validator.String{
					stringvalidator.OneOf("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag"),
				},
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port MTU.").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1500),
				Validators: []validator.Int64{
					int64validator.Between(576, 9216),
				},
			},
			"native_vlan": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("vlan-1"),
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port speed.").AddStringEnumDescription("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G").AddDefaultValueDescription("auto").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("auto"),
				Validators: []validator.String{
					stringvalidator.OneOf("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G"),
				},
			},
			"trunk_vlans": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of trunk VLANs.").AddDefaultValueDescription("1-4094").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("1-4094"),
			},
			"user_configured_flags": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port User Config Flags.").String,
				Optional:            true,
			},
		},
	}
}

func (r *PortChannelInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *PortChannelInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PortChannelInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	client, ok := r.clients[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	// Post object
	body := plan.toBody(false)
	_, err := client.Post(plan.getDn(), body.Str)
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

func (r *PortChannelInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PortChannelInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	client, ok := r.clients[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
	res, err := client.GetDn(state.Dn.ValueString(), queries...)
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

func (r *PortChannelInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PortChannelInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	client, ok := r.clients[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	body := plan.toBody(false)

	_, err := client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PortChannelInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PortChannelInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	client, ok := r.clients[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	body := state.toDeleteBody()

	if len(body.Str) > 0 {
		_, err := client.Post(state.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	} else {
		res, err := client.DeleteDn(state.Dn.ValueString())
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

func (r *PortChannelInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
