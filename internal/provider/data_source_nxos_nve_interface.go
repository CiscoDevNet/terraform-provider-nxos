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
	_ datasource.DataSource              = &NVEInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &NVEInterfaceDataSource{}
)

func NewNVEInterfaceDataSource() datasource.DataSource {
	return &NVEInterfaceDataSource{}
}

type NVEInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *NVEInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nve_interface"
}

func (d *NVEInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the NVE interface configuration.", "nvoEp", "Network%20Virtualization/nvo:Ep/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: "Administrative state.",
				Computed:            true,
			},
			"advertise_virtual_mac": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable Virtual MAC Advertisement in VPC mode.",
				Computed:            true,
			},
			"hold_down_time": schema.Int64Attribute{
				MarkdownDescription: "Hold Down Time.",
				Computed:            true,
			},
			"host_reachability_protocol": schema.StringAttribute{
				MarkdownDescription: "Host Reachability Protocol.",
				Computed:            true,
			},
			"ingress_replication_protocol_bgp": schema.BoolAttribute{
				MarkdownDescription: "VxLAN Ingress Replication Protocol BGP.",
				Computed:            true,
			},
			"multicast_group_l2": schema.StringAttribute{
				MarkdownDescription: "Base multicast group address for L2.",
				Computed:            true,
			},
			"multicast_group_l3": schema.StringAttribute{
				MarkdownDescription: "Base multicast group address for L3.",
				Computed:            true,
			},
			"multisite_source_interface": schema.StringAttribute{
				MarkdownDescription: "Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`.",
				Computed:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: "Source Interface associated with the NVE. Must match first field in the output of `show int brief`.",
				Computed:            true,
			},
			"suppress_arp": schema.BoolAttribute{
				MarkdownDescription: "Suppress ARP.",
				Computed:            true,
			},
			"suppress_mac_route": schema.BoolAttribute{
				MarkdownDescription: "Suppress MAC Route.",
				Computed:            true,
			},
		},
	}
}

func (d *NVEInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *NVEInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NVEInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	device, ok := d.data.Devices[config.Device.ValueString()]
	if !ok {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find device '%s' in provider configuration", config.Device.ValueString()))
		return
	}

	queries := []func(*nxos.Req){}
	res, err := device.Client.GetDn(config.getDn(), queries...)
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
