// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type dataSourceBridgeDomainType struct{}

func (t dataSourceBridgeDomainType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read a bridge domain.", "l2BD", "Layer%202/l2:BD/").String,

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"fabric_encap": {
				MarkdownDescription: "Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Type:                types.StringType,
				Required:            true,
			},
			"access_encap": {
				MarkdownDescription: "Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Type:                types.StringType,
				Computed:            true,
			},
			"name": {
				MarkdownDescription: "Bridge domain name.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceBridgeDomainType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBridgeDomain{
		provider: provider,
	}, diags
}

type dataSourceBridgeDomain struct {
	provider provider
}

func (d dataSourceBridgeDomain) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BridgeDomain

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn())

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
