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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &VPCKeepaliveResource{}
var _ resource.ResourceWithImportState = &VPCKeepaliveResource{}

func NewVPCKeepaliveResource() resource.Resource {
	return &VPCKeepaliveResource{}
}

type VPCKeepaliveResource struct {
	clients map[string]*nxos.Client
}

func (r *VPCKeepaliveResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_keepalive"
}

func (r *VPCKeepaliveResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the vPC keepalive configuration.", "vpcKeepalive", "System/vpc:Keepalive/").AddParents("vpc_domain").AddChildren("vpc_peerlink").String,

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
			"destination_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive destination address.").String,
				Required:            true,
			},
			"flush_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive flush timeout.").AddIntegerRangeDescription(3, 10).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 10),
				},
			},
			"interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive interval.").AddIntegerRangeDescription(400, 10000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(400, 10000),
				},
			},
			"precedence_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive precedence type. `0` - network, `1` - internet, `2` - critical, `3` flash-override, `4` - flash, `5` - immediate, `6` - prioriy, `7` - routine.").AddIntegerRangeDescription(0, 7).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"precedence_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive precedence value.").AddIntegerRangeDescription(0, 7).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"source_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive source address.").String,
				Required:            true,
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive timeout.").AddIntegerRangeDescription(3, 20).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(3, 20),
				},
			},
			"type_of_service_byte": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive type of service (ToS) byte.").AddIntegerRangeDescription(0, 255).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"type_of_service_configuration_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive type of service (ToS) configuration type. `0` - noCfg, `1` - tos-byte, `2` - tos-value, `3` - tos-type, `4` -  precedence-type, `5` - precedence-value.").AddIntegerRangeDescription(0, 5).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 5),
				},
			},
			"type_of_service_type": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive type of service (ToS) type. `0` - min-delay, `1` - max-throughput, `2` - max-reliability, `3` - min-monetary-cost, `4` -  normal.").AddIntegerRangeDescription(0, 4).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 4),
				},
			},
			"type_of_service_value": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive type of service (ToS) value.").AddIntegerRangeDescription(0, 15).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 15),
				},
			},
			"udp_port": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive UDP port.").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vPC Keepalive VRF.").String,
				Optional:            true,
			},
		},
	}
}

func (r *VPCKeepaliveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *VPCKeepaliveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VPCKeepalive

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.clients[plan.Device.ValueString()].Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VPCKeepaliveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VPCKeepalive

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

	state.fromBody(res, false)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *VPCKeepaliveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan VPCKeepalive

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.clients[plan.Device.ValueString()].Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VPCKeepaliveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VPCKeepalive

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	res, err := r.clients[state.Device.ValueString()].DeleteDn(state.Dn.ValueString())
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "1" && errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *VPCKeepaliveResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}