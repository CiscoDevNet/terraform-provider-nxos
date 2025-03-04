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
	_ datasource.DataSource              = &IPv4StaticRouteDataSource{}
	_ datasource.DataSourceWithConfigure = &IPv4StaticRouteDataSource{}
)

func NewIPv4StaticRouteDataSource() datasource.DataSource {
	return &IPv4StaticRouteDataSource{}
}

type IPv4StaticRouteDataSource struct {
	data *NxosProviderData
}

func (d *IPv4StaticRouteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv4_static_route"
}

func (d *IPv4StaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read an IPv4 static route.", "ipv4Route", "Layer%203/ipv4:Route/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"prefix": schema.StringAttribute{
				MarkdownDescription: "Prefix.",
				Required:            true,
			},
			"next_hops": schema.ListNestedAttribute{
				MarkdownDescription: "List of next hops.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: "Must match first field in the output of `show intf brief` or `unspecified`. Example: `eth1/1` or `vlan100`.",
							Computed:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: "Nexthop address.",
							Computed:            true,
						},
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "Nexthop VRF.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description.",
							Computed:            true,
						},
						"object": schema.Int64Attribute{
							MarkdownDescription: "Object to be tracked.",
							Computed:            true,
						},
						"preference": schema.Int64Attribute{
							MarkdownDescription: "Route preference.",
							Computed:            true,
						},
						"tag": schema.Int64Attribute{
							MarkdownDescription: "Tag value.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *IPv4StaticRouteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *IPv4StaticRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config IPv4StaticRoute

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
	queries = append(queries, nxos.Query("rsp-subtree", "children"))
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
