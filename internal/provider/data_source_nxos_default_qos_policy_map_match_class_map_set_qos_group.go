// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource{}

func NewDefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource() datasource.DataSource {
	return &DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource{}
}

type DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource struct {
	data *NxosProviderData
}

func (d *DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_qos_policy_map_match_class_map_set_qos_group"
}

func (d *DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the default QoS policy map match class map set QoS group configuration.", "ipqosSetQoSGrp", "Qos/ipqos:SetQoSGrp/").String,

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
			"policy_map_name": {
				MarkdownDescription: "Policy map name.",
				Type:                types.StringType,
				Required:            true,
			},
			"class_map_name": {
				MarkdownDescription: "Class map name.",
				Type:                types.StringType,
				Required:            true,
			},
			"qos_group_id": {
				MarkdownDescription: "QoS group ID.",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (d *DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*NxosProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected data, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.data = data
}

func (d *DefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state DefaultQOSPolicyMapMatchClassMapSetQOSGroup

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.data.client.GetDn(config.getDn(), nxos.OverrideUrl(d.data.devices[config.Device.Value]))

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
