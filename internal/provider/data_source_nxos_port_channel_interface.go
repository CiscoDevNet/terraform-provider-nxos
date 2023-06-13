// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &PortChannelInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &PortChannelInterfaceDataSource{}
)

func NewPortChannelInterfaceDataSource() datasource.DataSource {
	return &PortChannelInterfaceDataSource{}
}

type PortChannelInterfaceDataSource struct {
	clients map[string]*nxos.Client
}

func (d *PortChannelInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_port_channel_interface"
}

func (d *PortChannelInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the configuration of a port-channel interface.", "pcAggrIf", "Interfaces/pc:AggrIf/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `po1`.",
				Required:            true,
			},
			"port_channel_mode": schema.StringAttribute{
				MarkdownDescription: "The aggregated interface protocol channel mode.",
				Computed:            true,
			},
			"minimum_links": schema.Int64Attribute{
				MarkdownDescription: "Minimum links.",
				Computed:            true,
			},
			"maximum_links": schema.Int64Attribute{
				MarkdownDescription: "Maximum links.",
				Computed:            true,
			},
			"suspend_individual": schema.StringAttribute{
				MarkdownDescription: "Suspend Individual Port.",
				Computed:            true,
			},
			"access_vlan": schema.StringAttribute{
				MarkdownDescription: "Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Computed:            true,
			},
			"admin_state": schema.StringAttribute{
				MarkdownDescription: "Administrative port state.",
				Computed:            true,
			},
			"auto_negotiation": schema.StringAttribute{
				MarkdownDescription: "Administrative port auto-negotiation.",
				Computed:            true,
			},
			"bandwidth": schema.Int64Attribute{
				MarkdownDescription: "The bandwidth parameter for a routed interface, port channel, or subinterface.",
				Computed:            true,
			},
			"delay": schema.Int64Attribute{
				MarkdownDescription: "The administrative port delay time.",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Interface description.",
				Computed:            true,
			},
			"duplex": schema.StringAttribute{
				MarkdownDescription: "Duplex.",
				Computed:            true,
			},
			"layer": schema.StringAttribute{
				MarkdownDescription: "Administrative port layer.",
				Computed:            true,
			},
			"link_logging": schema.StringAttribute{
				MarkdownDescription: "Administrative link logging.",
				Computed:            true,
			},
			"medium": schema.StringAttribute{
				MarkdownDescription: "The administrative port medium type.",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Administrative port mode.",
				Computed:            true,
			},
			"mtu": schema.Int64Attribute{
				MarkdownDescription: "Administrative port MTU.",
				Computed:            true,
			},
			"native_vlan": schema.StringAttribute{
				MarkdownDescription: "Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.",
				Computed:            true,
			},
			"speed": schema.StringAttribute{
				MarkdownDescription: "Administrative port speed.",
				Computed:            true,
			},
			"trunk_vlans": schema.StringAttribute{
				MarkdownDescription: "List of trunk VLANs.",
				Computed:            true,
			},
			"user_configured_flags": schema.StringAttribute{
				MarkdownDescription: "Port User Config Flags.",
				Computed:            true,
			},
		},
	}
}

func (d *PortChannelInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *PortChannelInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config PortChannelInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	queries := []func(*nxos.Req){}
	res, err := d.clients[config.Device.ValueString()].GetDn(config.getDn(), queries...)
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
