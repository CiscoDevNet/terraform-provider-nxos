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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
	_ datasource.DataSourceWithConfigure = &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
)

func NewDefaultQOSPolicyMapMatchClassMapPoliceDataSource() datasource.DataSource {
	return &DefaultQOSPolicyMapMatchClassMapPoliceDataSource{}
}

type DefaultQOSPolicyMapMatchClassMapPoliceDataSource struct {
	clients map[string]*nxos.Client
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_qos_policy_map_match_class_map_police"
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the default QoS policy map match class map police configuration.", "ipqosPolice", "Qos/ipqos:Police/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"policy_map_name": schema.StringAttribute{
				MarkdownDescription: "Policy map name.",
				Required:            true,
			},
			"class_map_name": schema.StringAttribute{
				MarkdownDescription: "Class map name.",
				Required:            true,
			},
			"bc_rate": schema.Int64Attribute{
				MarkdownDescription: "CIR burst rate.",
				Computed:            true,
			},
			"bc_unit": schema.StringAttribute{
				MarkdownDescription: "CIR burst rate unit.",
				Computed:            true,
			},
			"be_rate": schema.Int64Attribute{
				MarkdownDescription: "PIR burst rate.",
				Computed:            true,
			},
			"be_unit": schema.StringAttribute{
				MarkdownDescription: "PIR burst rate unit.",
				Computed:            true,
			},
			"cir_rate": schema.Int64Attribute{
				MarkdownDescription: "CIR rate.",
				Computed:            true,
			},
			"cir_unit": schema.StringAttribute{
				MarkdownDescription: "CIR rate unit.",
				Computed:            true,
			},
			"conform_action": schema.StringAttribute{
				MarkdownDescription: "Conform action.",
				Computed:            true,
			},
			"conform_set_cos": schema.Int64Attribute{
				MarkdownDescription: "Set CoS for conforming traffic.",
				Computed:            true,
			},
			"conform_set_dscp": schema.Int64Attribute{
				MarkdownDescription: "Set DSCP for conforming traffic.",
				Computed:            true,
			},
			"conform_set_precedence": schema.StringAttribute{
				MarkdownDescription: "Set precedence for conforming traffic.",
				Computed:            true,
			},
			"conform_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: "Set qos-group for conforming traffic.",
				Computed:            true,
			},
			"exceed_action": schema.StringAttribute{
				MarkdownDescription: "Exceed action.",
				Computed:            true,
			},
			"exceed_set_cos": schema.Int64Attribute{
				MarkdownDescription: "Set CoS for exceeding traffic.",
				Computed:            true,
			},
			"exceed_set_dscp": schema.Int64Attribute{
				MarkdownDescription: "Set DSCP for exceeding traffic.",
				Computed:            true,
			},
			"exceed_set_precedence": schema.StringAttribute{
				MarkdownDescription: "Set precedence for exceeding traffic.",
				Computed:            true,
			},
			"exceed_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: "Set qos-group for exceeding traffic.",
				Computed:            true,
			},
			"pir_rate": schema.Int64Attribute{
				MarkdownDescription: "PIR rate.",
				Computed:            true,
			},
			"pir_unit": schema.StringAttribute{
				MarkdownDescription: "PIR rate unit.",
				Computed:            true,
			},
			"violate_action": schema.StringAttribute{
				MarkdownDescription: "Violate action.",
				Computed:            true,
			},
			"violate_set_cos": schema.Int64Attribute{
				MarkdownDescription: "Set CoS for violating traffic.",
				Computed:            true,
			},
			"violate_set_dscp": schema.Int64Attribute{
				MarkdownDescription: "Set DSCP for violating traffic.",
				Computed:            true,
			},
			"violate_set_precedence": schema.StringAttribute{
				MarkdownDescription: "Set precedence for violating traffic.",
				Computed:            true,
			},
			"violate_set_qos_group": schema.Int64Attribute{
				MarkdownDescription: "Set qos-group for violating traffic.",
				Computed:            true,
			},
		},
	}
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *DefaultQOSPolicyMapMatchClassMapPoliceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DefaultQOSPolicyMapMatchClassMapPolice

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	client, ok := d.clients[config.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", config.Device.ValueString()))
		return
	}

	queries := []func(*nxos.Req){}
	res, err := client.GetDn(config.getDn(), queries...)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(res, true)
	config.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
