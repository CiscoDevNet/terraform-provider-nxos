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
var _ resource.Resource = &DefaultQOSPolicyMapMatchClassMapPoliceResource{}
var _ resource.ResourceWithImportState = &DefaultQOSPolicyMapMatchClassMapPoliceResource{}

func NewDefaultQOSPolicyMapMatchClassMapPoliceResource() resource.Resource {
	return &DefaultQOSPolicyMapMatchClassMapPoliceResource{}
}

type DefaultQOSPolicyMapMatchClassMapPoliceResource struct {
	data *NxosProviderData
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_qos_policy_map_match_class_map_police"
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the default QoS policy map match class map police configuration.", "ipqosPolice", "Qos/ipqos:Police/").AddParents("default_qos_policy_map_match_class_map").String,

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
			"policy_map_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy map name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"class_map_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Class map name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"bc_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("CIR burst rate.").AddIntegerRangeDescription(0, 536870912).AddDefaultValueDescription("200").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(200),
				Validators: []validator.Int64{
					int64validator.Between(0, 536870912),
				},
			},
			"bc_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CIR burst rate unit.").AddStringEnumDescription("unspecified", "bytes", "kbytes", "mbytes", "ms", "us", "packets").AddDefaultValueDescription("ms").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ms"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "bytes", "kbytes", "mbytes", "ms", "us", "packets"),
				},
			},
			"be_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("PIR burst rate.").AddIntegerRangeDescription(0, 536870912).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 536870912),
				},
			},
			"be_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PIR burst rate unit.").AddStringEnumDescription("unspecified", "bytes", "kbytes", "mbytes", "ms", "us", "packets").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "bytes", "kbytes", "mbytes", "ms", "us", "packets"),
				},
			},
			"cir_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("CIR rate.").AddIntegerRangeDescription(0, 100000000000).AddDefaultValueDescription("0").String,
				Required:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100000000000),
				},
			},
			"cir_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CIR rate unit.").AddStringEnumDescription("unspecified", "bps", "kbps", "mbps", "gbps", "pps", "pct").AddDefaultValueDescription("bps").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("bps"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "bps", "kbps", "mbps", "gbps", "pps", "pct"),
				},
			},
			"conform_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Conform action.").AddStringEnumDescription("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit").AddDefaultValueDescription("transmit").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("transmit"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit"),
				},
			},
			"conform_set_cos": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set CoS for conforming traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"conform_set_dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set DSCP for conforming traffic.").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"conform_set_precedence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set precedence for conforming traffic.").AddStringEnumDescription("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network").AddDefaultValueDescription("routine").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("routine"),
				Validators: []validator.String{
					stringvalidator.OneOf("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network"),
				},
			},
			"conform_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set qos-group for conforming traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"exceed_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Exceed action.").AddStringEnumDescription("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit"),
				},
			},
			"exceed_set_cos": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set CoS for exceeding traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"exceed_set_dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set DSCP for exceeding traffic.").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"exceed_set_precedence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set precedence for exceeding traffic.").AddStringEnumDescription("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network").AddDefaultValueDescription("routine").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("routine"),
				Validators: []validator.String{
					stringvalidator.OneOf("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network"),
				},
			},
			"exceed_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set qos-group for exceeding traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"pir_rate": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("PIR rate.").AddIntegerRangeDescription(0, 100000000000).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 100000000000),
				},
			},
			"pir_unit": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("PIR rate unit.").AddStringEnumDescription("unspecified", "bps", "kbps", "mbps", "gbps", "pps", "pct").AddDefaultValueDescription("unspecified").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("unspecified"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "bps", "kbps", "mbps", "gbps", "pps", "pct"),
				},
			},
			"violate_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Violate action.").AddStringEnumDescription("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit").AddDefaultValueDescription("drop").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("drop"),
				Validators: []validator.String{
					stringvalidator.OneOf("unspecified", "transmit", "drop", "set-cos-transmit", "set-dscp-transmit", "set-prec-transmit", "set-qos-transmit"),
				},
			},
			"violate_set_cos": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set CoS for violating traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
			"violate_set_dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set DSCP for violating traffic.").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"violate_set_precedence": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set precedence for violating traffic.").AddStringEnumDescription("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network").AddDefaultValueDescription("routine").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("routine"),
				Validators: []validator.String{
					stringvalidator.OneOf("routine", "priority", "immediate", "flash", "flash-override", "critical", "internet", "network"),
				},
			},
			"violate_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set qos-group for violating traffic.").AddIntegerRangeDescription(0, 7).AddDefaultValueDescription("0").String,
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				Validators: []validator.Int64{
					int64validator.Between(0, 7),
				},
			},
		},
	}
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DefaultQOSPolicyMapMatchClassMapPolice

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	// Post object
	if device.Managed {
		body := plan.toBody(false)
		_, err := device.Client.Post(plan.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
			return
		}
	}

	plan.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DefaultQOSPolicyMapMatchClassMapPolice

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		queries := []func(*nxos.Req){nxos.Query("rsp-prop-include", "config-only")}
		res, err := device.Client.GetDn(state.Dn.ValueString(), queries...)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}

		imp, diags := helpers.IsFlagImporting(ctx, req)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromBody(res, imp)
		if imp {
			state.getIdsFromDn()
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DefaultQOSPolicyMapMatchClassMapPolice

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	device, ok := r.data.Devices[plan.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", plan.Device.ValueString()))
		return
	}

	if device.Managed {

		body := plan.toBody(false)

		_, err := device.Client.Post(plan.getDn(), body.Str)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DefaultQOSPolicyMapMatchClassMapPolice

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	device, ok := r.data.Devices[state.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", state.Device.ValueString()))
		return
	}

	if device.Managed {
		body := state.toDeleteBody()

		if len(body.Str) > 0 {
			_, err := device.Client.Post(state.getDn(), body.Str)
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
				return
			}
		} else {
			res, err := device.Client.DeleteDn(state.Dn.ValueString())
			if err != nil {
				errCode := res.Get("imdata.0.error.attributes.code").Str
				// Ignore errors of type "Cannot delete object"
				if errCode != "1" && errCode != "107" {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object, got error: %s", err))
					return
				}
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *DefaultQOSPolicyMapMatchClassMapPoliceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
