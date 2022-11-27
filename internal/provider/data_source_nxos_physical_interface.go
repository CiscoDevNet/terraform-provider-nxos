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
	_ datasource.DataSource              = &PhysicalInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &PhysicalInterfaceDataSource{}
)

func NewPhysicalInterfaceDataSource() datasource.DataSource {
	return &PhysicalInterfaceDataSource{}
}

type PhysicalInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *PhysicalInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_physical_interface"
}

func (d *PhysicalInterfaceDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the configuration of a physical interface.", "l1PhysIf", "System/l1:PhysIf/").String,

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
			"fec_mode": {
				MarkdownDescription: "FEC mode.",
				Type:                types.StringType,
				Computed:            true,
			},
			"access_vlan": {
				MarkdownDescription: "Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Type:                types.StringType,
				Computed:            true,
			},
			"admin_state": {
				MarkdownDescription: "Administrative port state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"auto_negotiation": {
				MarkdownDescription: "Administrative port auto-negotiation.",
				Type:                types.StringType,
				Computed:            true,
			},
			"bandwidth": {
				MarkdownDescription: "The bandwidth parameter for a routed interface, port channel, or subinterface.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"delay": {
				MarkdownDescription: "The administrative port delay time.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: "Interface description.",
				Type:                types.StringType,
				Computed:            true,
			},
			"duplex": {
				MarkdownDescription: "Duplex.",
				Type:                types.StringType,
				Computed:            true,
			},
			"layer": {
				MarkdownDescription: "Administrative port layer.",
				Type:                types.StringType,
				Computed:            true,
			},
			"link_logging": {
				MarkdownDescription: "Administrative link logging.",
				Type:                types.StringType,
				Computed:            true,
			},
			"link_debounce_down": {
				MarkdownDescription: "Administrative port link debounce interval.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"link_debounce_up": {
				MarkdownDescription: "Link Debounce Interval - LinkUp Event.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"medium": {
				MarkdownDescription: "The administrative port medium type.",
				Type:                types.StringType,
				Computed:            true,
			},
			"mode": {
				MarkdownDescription: "Administrative port mode.",
				Type:                types.StringType,
				Computed:            true,
			},
			"mtu": {
				MarkdownDescription: "Administrative port MTU.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"native_vlan": {
				MarkdownDescription: "Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Type:                types.StringType,
				Computed:            true,
			},
			"speed": {
				MarkdownDescription: "Administrative port speed.",
				Type:                types.StringType,
				Computed:            true,
			},
			"speed_group": {
				MarkdownDescription: "Speed group.",
				Type:                types.StringType,
				Computed:            true,
			},
			"trunk_vlans": {
				MarkdownDescription: "List of trunk VLANs.",
				Type:                types.StringType,
				Computed:            true,
			},
			"uni_directional_ethernet": {
				MarkdownDescription: "UDE (Uni-Directional Ethernet).",
				Type:                types.StringType,
				Computed:            true,
			},
			"user_configured_flags": {
				MarkdownDescription: "Port User Config Flags.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (d *PhysicalInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *PhysicalInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state PhysicalInterface

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
