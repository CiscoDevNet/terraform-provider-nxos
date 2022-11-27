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

func (d *NVEInterfaceDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the NVE interface configuration.", "nvoEp", "Network%20Virtualization/nvo:Ep/").String,

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
			"admin_state": {
				MarkdownDescription: "Administrative state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"advertise_virtual_mac": {
				MarkdownDescription: "Enable or disable Virtual MAC Advertisement in VPC mode.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"hold_down_time": {
				MarkdownDescription: "Hold Down Time.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"host_reachability_protocol": {
				MarkdownDescription: "Host Reachability Protocol.",
				Type:                types.StringType,
				Computed:            true,
			},
			"ingress_replication_protocol_bgp": {
				MarkdownDescription: "VxLAN Ingress Replication Protocol BGP.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"multicast_group_l2": {
				MarkdownDescription: "Base multicast group address for L2.",
				Type:                types.StringType,
				Computed:            true,
			},
			"multicast_group_l3": {
				MarkdownDescription: "Base multicast group address for L3.",
				Type:                types.StringType,
				Computed:            true,
			},
			"multisite_source_interface": {
				MarkdownDescription: "Interface representing the Multisite Border Gateway. Must match first field in the output of `show int brief`.",
				Type:                types.StringType,
				Computed:            true,
			},
			"source_interface": {
				MarkdownDescription: "Source Interface associated with the NVE. Must match first field in the output of `show int brief`.",
				Type:                types.StringType,
				Computed:            true,
			},
			"suppress_arp": {
				MarkdownDescription: "Suppress ARP.",
				Type:                types.BoolType,
				Computed:            true,
			},
			"suppress_mac_route": {
				MarkdownDescription: "Suppress MAC Route.",
				Type:                types.BoolType,
				Computed:            true,
			},
		},
	}, nil
}

func (d *NVEInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *NVEInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state NVEInterface

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
	state.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
