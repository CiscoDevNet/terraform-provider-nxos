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
	_ datasource.DataSource              = &BGPPeerTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &BGPPeerTemplateDataSource{}
)

func NewBGPPeerTemplateDataSource() datasource.DataSource {
	return &BGPPeerTemplateDataSource{}
}

type BGPPeerTemplateDataSource struct {
	clients map[string]*nxos.Client
}

func (d *BGPPeerTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bgp_peer_template"
}

func (d *BGPPeerTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer template configuration.", "bgpPeerCont", "Routing%20and%20Forwarding/bgp:PeerCont/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"asn": schema.StringAttribute{
				MarkdownDescription: "Autonomous system number.",
				Required:            true,
			},
			"template_name": schema.StringAttribute{
				MarkdownDescription: "Peer template name.",
				Required:            true,
			},
			"remote_asn": schema.StringAttribute{
				MarkdownDescription: "Peer template autonomous system number.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Peer template description.",
				Computed:            true,
			},
			"peer_type": schema.StringAttribute{
				MarkdownDescription: "Neighbor Fabric Type.",
				Computed:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: "Source Interface. Must match first field in the output of `show intf brief`.",
				Computed:            true,
			},
		},
	}
}

func (d *BGPPeerTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *BGPPeerTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config BGPPeerTemplate

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
