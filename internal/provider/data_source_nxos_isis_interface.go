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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ISISInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &ISISInterfaceDataSource{}
)

func NewISISInterfaceDataSource() datasource.DataSource {
	return &ISISInterfaceDataSource{}
}

type ISISInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *ISISInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_interface"
}

func (d *ISISInterfaceDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the IS-IS interface configuration.", "isisInternalIf", "Routing%20and%20Forwarding/isis:InternalIf/").String,

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
			"interface_id": {
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Type:                types.StringType,
				Required:            true,
			},
			"authentication_check": {
				MarkdownDescription: "Authentication Check for ISIS without specific level.",
				Type:                types.BoolType,
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
			"authentication_key": {
				MarkdownDescription: "Authentication Key for IS-IS without specific level.",
				Type:                types.StringType,
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
			"authentication_type": {
				MarkdownDescription: "IS-IS Authentication-Type without specific level.",
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
			"circuit_type": {
				MarkdownDescription: "Circuit type.",
				Type:                types.StringType,
				Computed:            true,
			},
			"vrf": {
				MarkdownDescription: "VRF.",
				Type:                types.StringType,
				Computed:            true,
			},
			"hello_interval": {
				MarkdownDescription: "Hello interval.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_interval_l1": {
				MarkdownDescription: "Hello interval Level-1.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_interval_l2": {
				MarkdownDescription: "Hello interval Level-2.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_multiplier": {
				MarkdownDescription: "Hello multiplier.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_multiplier_l1": {
				MarkdownDescription: "Hello multiplier Level-1.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_multiplier_l2": {
				MarkdownDescription: "Hello multiplier Level-2.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"hello_padding": {
				MarkdownDescription: "Hello padding.",
				Type:                types.StringType,
				Computed:            true,
			},
			"metric_l1": {
				MarkdownDescription: "Interface metric Level-1.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"metric_l2": {
				MarkdownDescription: "Interface metric Level-2.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"mtu_check": {
				MarkdownDescription: "MTU Check for IS-IS without specific level.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mtu_check_l1": {
				MarkdownDescription: "MTU Check for IS-IS on Level-1.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"mtu_check_l2": {
				MarkdownDescription: "MTU Check for IS-IS on Level-2.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"network_type_p2p": {
				MarkdownDescription: "Enabling Point-to-Point Network Type on IS-IS Interface.",
				Type:                types.StringType,
				Computed:            true,
			},
			"passive": {
				MarkdownDescription: "IS-IS Passive Interface Info.",
				Type:                types.StringType,
				Computed:            true,
			},
			"priority_l1": {
				MarkdownDescription: "Circuit priority.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"priority_l2": {
				MarkdownDescription: "Circuit priority.",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (d *ISISInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *ISISInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state ISISInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.data.client.GetDn(config.getDn(), nxos.OverrideUrl(d.data.devices[config.Device.ValueString()]))

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
