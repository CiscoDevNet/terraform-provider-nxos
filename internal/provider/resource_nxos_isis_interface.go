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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ISISInterfaceResource{}
var _ resource.ResourceWithImportState = &ISISInterfaceResource{}

func NewISISInterfaceResource() resource.Resource {
	return &ISISInterfaceResource{}
}

type ISISInterfaceResource struct {
	clients map[string]*nxos.Client
}

func (r *ISISInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_interface"
}

func (r *ISISInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the IS-IS interface configuration.", "isisInternalIf", "Routing%20and%20Forwarding/isis:InternalIf/").AddParents("isis").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"authentication_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS without specific level.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"authentication_check_l1": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-1.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"authentication_check_l2": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-2.").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"authentication_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS without specific level.").String,
				Optional:            true,
			},
			"authentication_key_l1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-1.").String,
				Optional:            true,
			},
			"authentication_key_l2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-2.").String,
				Optional:            true,
			},
			"authentication_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type without specific level.").AddStringEnumDescription("clear", "md5", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "md5", "unknown"),
				},
			},
			"authentication_type_l1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-1.").AddStringEnumDescription("clear", "md5", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "md5", "unknown"),
				},
			},
			"authentication_type_l2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-2.").AddStringEnumDescription("clear", "md5", "unknown").AddDefaultValueDescription("unknown").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unknown"),
				Validators: []validator.String{
					stringvalidator.OneOf("clear", "md5", "unknown"),
				},
			},
			"circuit_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Circuit type.").AddStringEnumDescription("l1", "l2", "l12").AddDefaultValueDescription("l12").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("l12"),
				Validators: []validator.String{
					stringvalidator.OneOf("l1", "l2", "l12"),
				},
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF.").String,
				Optional:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"hello_interval_l1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval Level-1.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"hello_interval_l2": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello interval Level-2.").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10),
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"hello_multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello multiplier.").AddIntegerRangeDescription(3, 1000).AddDefaultValueDescription("3").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
				Validators: []validator.Int64{
					int64validator.Between(3, 1000),
				},
			},
			"hello_multiplier_l1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello multiplier Level-1.").AddIntegerRangeDescription(3, 1000).AddDefaultValueDescription("3").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
				Validators: []validator.Int64{
					int64validator.Between(3, 1000),
				},
			},
			"hello_multiplier_l2": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello multiplier Level-2.").AddIntegerRangeDescription(3, 1000).AddDefaultValueDescription("3").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(3),
				Validators: []validator.Int64{
					int64validator.Between(3, 1000),
				},
			},
			"hello_padding": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hello padding.").AddStringEnumDescription("always", "transient", "never").AddDefaultValueDescription("always").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("always"),
				Validators: []validator.String{
					stringvalidator.OneOf("always", "transient", "never"),
				},
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Instance to which the interface belongs to.").String,
				Optional:            true,
			},
			"metric_l1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface metric Level-1.").AddIntegerRangeDescription(0, 16777216).AddDefaultValueDescription("16777216").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(16777216),
				Validators: []validator.Int64{
					int64validator.Between(0, 16777216),
				},
			},
			"metric_l2": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface metric Level-2.").AddIntegerRangeDescription(0, 16777216).AddDefaultValueDescription("16777216").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(16777216),
				Validators: []validator.Int64{
					int64validator.Between(0, 16777216),
				},
			},
			"mtu_check": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MTU Check for IS-IS without specific level.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mtu_check_l1": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MTU Check for IS-IS on Level-1.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mtu_check_l2": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MTU Check for IS-IS on Level-2.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"network_type_p2p": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enabling Point-to-Point Network Type on IS-IS Interface.").AddStringEnumDescription("off", "on", "useAllISMac").AddDefaultValueDescription("off").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("off"),
				Validators: []validator.String{
					stringvalidator.OneOf("off", "on", "useAllISMac"),
				},
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Passive Interface Info.").AddStringEnumDescription("l1", "l2", "l12", "noL1", "noL2", "noL12", "inheritDef").AddDefaultValueDescription("inheritDef").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("inheritDef"),
				Validators: []validator.String{
					stringvalidator.OneOf("l1", "l2", "l12", "noL1", "noL2", "noL12", "inheritDef"),
				},
			},
			"priority_l1": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Circuit priority.").AddIntegerRangeDescription(0, 127).AddDefaultValueDescription("64").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(64),
				Validators: []validator.Int64{
					int64validator.Between(0, 127),
				},
			},
			"priority_l2": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Circuit priority.").AddIntegerRangeDescription(0, 127).AddDefaultValueDescription("64").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(64),
				Validators: []validator.Int64{
					int64validator.Between(0, 127),
				},
			},
			"enable_ipv4": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enabling ISIS router tag on Interface's IPV4 family.").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}

func (r *ISISInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *ISISInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ISISInterface

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

func (r *ISISInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ISISInterface

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

func (r *ISISInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ISISInterface

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

func (r *ISISInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ISISInterface

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
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
				return
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *ISISInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
