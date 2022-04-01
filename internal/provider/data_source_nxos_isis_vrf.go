// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type dataSourceISISVRFType struct{}

func (t dataSourceISISVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the IS-IS VRF configuration.", "isisDom", "Routing%20and%20Forwarding/isis:Dom/").String,

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"instance_name": {
				MarkdownDescription: "IS-IS instance name.",
				Type:                types.StringType,
				Required:            true,
			},
			"name": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"admin_state": {
				MarkdownDescription: "Administrative state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"authentication_check_l1": {
				MarkdownDescription: "Authentication Check for ISIS on Level-1.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"authentication_check_l2": {
				MarkdownDescription: "Authentication Check for ISIS on Level-2.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"authentication_key_l1": {
				MarkdownDescription: "Authentication Key for IS-IS on Level-1.",
				Type:                types.StringType,
				Computed:            true,
			},
			"authentication_key_l2": {
				MarkdownDescription: "Authentication Key for IS-IS on Level-2.",
				Type:                types.StringType,
				Computed:            true,
			},
			"authentication_type_l1": {
				MarkdownDescription: "IS-IS Authentication-Type for Level-1.",
				Type:                types.StringType,
				Computed:            true,
			},
			"authentication_type_l2": {
				MarkdownDescription: "IS-IS Authentication-Type for Level-2.",
				Type:                types.StringType,
				Computed:            true,
			},
			"bandwidth_reference": {
				MarkdownDescription: "The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost metric.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"banwidth_reference_unit": {
				MarkdownDescription: "Bandwidth reference unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"is_type": {
				MarkdownDescription: "IS-IS domain type.",
				Type:                types.StringType,
				Computed:            true,
			},
			"metric_type": {
				MarkdownDescription: "IS-IS metric type.",
				Type:                types.StringType,
				Computed:            true,
			},
			"mtu": {
				MarkdownDescription: "The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352 bytes.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"net": {
				MarkdownDescription: "Holds IS-IS domain NET (address) value.",
				Type:                types.StringType,
				Computed:            true,
			},
			"passive_default": {
				MarkdownDescription: "IS-IS Domain passive-interface default level.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceISISVRFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceISISVRF{
		provider: provider,
	}, diags
}

type dataSourceISISVRF struct {
	provider provider
}

func (d dataSourceISISVRF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state ISISVRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn(), nxos.OverrideUrl(d.provider.devices[config.Device.Value]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}