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
	_ datasource.DataSource              = &ISISVRFDataSource{}
	_ datasource.DataSourceWithConfigure = &ISISVRFDataSource{}
)

func NewISISVRFDataSource() datasource.DataSource {
	return &ISISVRFDataSource{}
}

type ISISVRFDataSource struct {
	data *NxosProviderData
}

func (d *ISISVRFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_vrf"
}

func (d *ISISVRFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the IS-IS VRF configuration.", "isisDom", "Routing%20and%20Forwarding/isis:Dom/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"instance_name": schema.StringAttribute{
				MarkdownDescription: "IS-IS instance name.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "VRF name.",
				Required:            true,
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: "Administrative state.",
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
			"authentication_key_l1": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-1.",
				Computed:            true,
			},
			"authentication_key_l2": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-2.",
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
			"bandwidth_reference": schema.Int64Attribute{
				MarkdownDescription: "The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost metric.",
				Computed:            true,
			},
			"banwidth_reference_unit": schema.StringAttribute{
				MarkdownDescription: "Bandwidth reference unit.",
				Computed:            true,
			},
			"is_type": schema.StringAttribute{
				MarkdownDescription: "IS-IS domain type.",
				Computed:            true,
			},
			"metric_type": schema.StringAttribute{
				MarkdownDescription: "IS-IS metric type.",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352 bytes.",
				Computed:            true,
			},
			"net": schema.StringAttribute{
				MarkdownDescription: "Holds IS-IS domain NET (address) value.",
				Computed:            true,
			},
			"passive_default": schema.StringAttribute{
				MarkdownDescription: "IS-IS Domain passive-interface default level.",
				Computed:            true,
			},
		},
	}
}

func (d *ISISVRFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *ISISVRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ISISVRF

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
