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

type dataSourceVRFRouteTargetType struct{}

func (t dataSourceVRFRouteTargetType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read a VRF Route Target Entry.", "rtctrlRttEntry", "Routing%20and%20Forwarding/rtctrl:RttEntry/").String,

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
			"vrf": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"address_family": {
				MarkdownDescription: "Address family.",
				Type:                types.StringType,
				Required:            true,
			},
			"route_target_address_family": {
				MarkdownDescription: "Route Target Address Family.",
				Type:                types.StringType,
				Required:            true,
			},
			"direction": {
				MarkdownDescription: "Route Target direction.",
				Type:                types.StringType,
				Required:            true,
			},
			"route_target": {
				MarkdownDescription: "Route Target in NX-OS DME format.",
				Type:                types.StringType,
				Required:            true,
			},
		},
	}, nil
}

func (t dataSourceVRFRouteTargetType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceVRFRouteTarget{
		provider: provider,
	}, diags
}

type dataSourceVRFRouteTarget struct {
	provider provider
}

func (d dataSourceVRFRouteTarget) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state VRFRouteTarget

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
