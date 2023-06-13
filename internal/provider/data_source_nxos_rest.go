package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RestDataSource{}
	_ datasource.DataSourceWithConfigure = &RestDataSource{}
)

func NewRestDataSource() datasource.DataSource {
	return &RestDataSource{}
}

type RestDataSource struct {
	clients map[string]*nxos.Client
}

func (d *RestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rest"
}

func (d *RestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read one NX-OS DME object.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"dn": schema.StringAttribute{
				MarkdownDescription: "Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].",
				Required:            true,
			},
			"class_name": schema.StringAttribute{
				MarkdownDescription: "Class name of object being retrieved.",
				Computed:            true,
			},
			"content": schema.MapAttribute{
				MarkdownDescription: "Map of key-value pairs which represents the attributes of object being retrieved.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *RestDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.clients = req.ProviderData.(map[string]*nxos.Client)
}

func (d *RestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state RestDataSourceModel

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.ValueString()))

	res, err := d.clients[config.Device.ValueString()].GetDn(config.Dn.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Check if we received an empty response without errors -> object has been deleted
	if !res.Exists() && err == nil {
		state.Id = types.StringNull()
	} else {
		// Set class_name
		var className string
		for class := range res.Map() {
			className = class
		}

		state.ClassName = types.StringValue(className)
		state.Dn = config.Dn
		state.Device = config.Device

		// Set content
		content := make(map[string]attr.Value)

		for attr, value := range res.Get(className + ".attributes").Map() {
			content[attr] = types.StringValue(value.String())
		}
		state.Content = types.MapValueMust(types.StringType, content)

		// Set id
		state.Id = state.Dn
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
