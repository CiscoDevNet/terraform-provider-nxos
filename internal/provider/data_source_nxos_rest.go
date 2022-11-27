package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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
	data *NxosProviderData
}

func (d *RestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rest"
}

func (d *RestDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read one NX-OS DME object.",

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
			"dn": {
				MarkdownDescription: "Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].",
				Type:                types.StringType,
				Required:            true,
			},
			"class_name": {
				MarkdownDescription: "Class name of object being retrieved.",
				Type:                types.StringType,
				Computed:            true,
			},
			"content": {
				MarkdownDescription: "Map of key-value pairs which represents the attributes of object being retrieved.",
				Type:                types.MapType{ElemType: types.StringType},
				Computed:            true,
			},
		},
	}, nil
}

func (d *RestDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
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

	res, err := d.data.client.GetDn(config.Dn.ValueString(), nxos.OverrideUrl(d.data.devices[config.Device.ValueString()]))
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
