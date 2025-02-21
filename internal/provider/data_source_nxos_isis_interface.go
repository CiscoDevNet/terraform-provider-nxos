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
	_ datasource.DataSource              = &ISISInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &ISISInterfaceDataSource{}
)

func NewISISInterfaceDataSource() datasource.DataSource {
	return &ISISInterfaceDataSource{}
}

type ISISInterfaceDataSource struct {
	clients map[string]*nxos.Client
}

func (d *ISISInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_interface"
}

func (d *ISISInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the IS-IS interface configuration.", "isisInternalIf", "Routing%20and%20Forwarding/isis:InternalIf/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Required:            true,
			},
			"authentication_check": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS without specific level.",
				Computed:            true,
			},
			"authentication_check_l1": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS on Level-1.",
				Computed:            true,
			},
			"authentication_check_l2": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS on Level-2.",
				Computed:            true,
			},
			"authentication_key": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS without specific level.",
				Computed:            true,
			},
			"authentication_key_l1": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-1.",
				Computed:            true,
			},
			"authentication_key_l2": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-2.",
				Computed:            true,
			},
			"authentication_type": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type without specific level.",
				Computed:            true,
			},
			"authentication_type_l1": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type for Level-1.",
				Computed:            true,
			},
			"authentication_type_l2": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type for Level-2.",
				Computed:            true,
			},
			"circuit_type": schema.StringAttribute{
				MarkdownDescription: "Circuit type.",
				Computed:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "VRF.",
				Computed:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Hello interval.",
				Computed:            true,
			},
			"hello_interval_l1": schema.Int64Attribute{
				MarkdownDescription: "Hello interval Level-1.",
				Computed:            true,
			},
			"hello_interval_l2": schema.Int64Attribute{
				MarkdownDescription: "Hello interval Level-2.",
				Computed:            true,
			},
			"hello_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier.",
				Computed:            true,
			},
			"hello_multiplier_l1": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier Level-1.",
				Computed:            true,
			},
			"hello_multiplier_l2": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier Level-2.",
				Computed:            true,
			},
			"hello_padding": schema.StringAttribute{
				MarkdownDescription: "Hello padding.",
				Computed:            true,
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: "Instance to which the interface belongs to.",
				Computed:            true,
			},
			"metric_l1": schema.Int64Attribute{
				MarkdownDescription: "Interface metric Level-1.",
				Computed:            true,
			},
			"metric_l2": schema.Int64Attribute{
				MarkdownDescription: "Interface metric Level-2.",
				Computed:            true,
			},
			"mtu_check": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS without specific level.",
				Computed:            true,
			},
			"mtu_check_l1": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS on Level-1.",
				Computed:            true,
			},
			"mtu_check_l2": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS on Level-2.",
				Computed:            true,
			},
			"network_type_p2p": schema.StringAttribute{
				MarkdownDescription: "Enabling Point-to-Point Network Type on IS-IS Interface.",
				Computed:            true,
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: "IS-IS Passive Interface Info.",
				Computed:            true,
			},
			"priority_l1": schema.Int64Attribute{
				MarkdownDescription: "Circuit priority.",
				Computed:            true,
			},
			"priority_l2": schema.Int64Attribute{
				MarkdownDescription: "Circuit priority.",
				Computed:            true,
			},
			"enable_ipv4": schema.BoolAttribute{
				MarkdownDescription: "Enabling ISIS router tag on Interface's IPV4 family.",
				Computed:            true,
			},
		},
	}
}

func (d *ISISInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *ISISInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ISISInterface

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
