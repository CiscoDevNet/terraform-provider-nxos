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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &BGPPeerTemplateAddressFamilyResource{}
var _ resource.ResourceWithImportState = &BGPPeerTemplateAddressFamilyResource{}

func NewBGPPeerTemplateAddressFamilyResource() resource.Resource {
	return &BGPPeerTemplateAddressFamilyResource{}
}

type BGPPeerTemplateAddressFamilyResource struct {
	clients map[string]*nxos.Client
}

func (r *BGPPeerTemplateAddressFamilyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_peer_template_address_family"
}

func (r *BGPPeerTemplateAddressFamilyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the BGP peer template address family configuration.", "bgpPeerAf", "Routing%20and%20Forwarding/bgp:PeerAf/").AddParents("bgp_peer_template").AddChildren("bgp_peer_template_max_prefix").String,

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
			"asn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Autonomous system number.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"template_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer template name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"address_family": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Address Family.").AddStringEnumDescription("ipv4-ucast", "vpnv4-ucast", "ipv6-ucast", "vpnv6-ucast", "l2vpn-evpn", "lnkstate").AddDefaultValueDescription("ipv4-ucast").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ipv4-ucast", "vpnv4-ucast", "ipv6-ucast", "vpnv6-ucast", "l2vpn-evpn", "lnkstate"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"control": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`, `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats:\n  - Single value. Example: `nh-self`\n  - Multiple values (comma-separated). Example: `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.").String,
				Optional:            true,
			},
			"send_community_extended": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Send-community extended.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
			"send_community_standard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Send-community standard.").AddStringEnumDescription("enabled", "disabled").AddDefaultValueDescription("disabled").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("disabled"),
				Validators: []validator.String{
					stringvalidator.OneOf("enabled", "disabled"),
				},
			},
		},
	}
}

func (r *BGPPeerTemplateAddressFamilyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (r *BGPPeerTemplateAddressFamilyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BGPPeerTemplateAddressFamily

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

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *BGPPeerTemplateAddressFamilyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state BGPPeerTemplateAddressFamily

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

func (r *BGPPeerTemplateAddressFamilyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BGPPeerTemplateAddressFamily

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

func (r *BGPPeerTemplateAddressFamilyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state BGPPeerTemplateAddressFamily

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

func (r *BGPPeerTemplateAddressFamilyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
