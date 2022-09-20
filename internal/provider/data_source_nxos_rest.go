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

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &RestDataSource{}

func NewRestDataSource() datasource.DataSource {
	return &RestDataSource{}
}

type RestDataSource struct {
	data *NxosProviderData
}

func (d *RestDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
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

func (d *RestDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *RestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state RestDataSourceModel

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.Value))

	res, err := d.data.client.GetDn(config.Dn.Value, nxos.OverrideUrl(d.data.devices[config.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	// Check if we received an empty response without errors -> object has been deleted
	if !res.Exists() && err == nil {
		state.Id.Value = ""
	} else {
		// Set class_name
		var className string
		for class := range res.Map() {
			className = class
		}

		state.ClassName.Value = className
		state.Dn.Value = config.Dn.Value
		state.Device = config.Device

		// Set content
		content := make(map[string]attr.Value)

		for attr, value := range res.Get(className + ".attributes").Map() {
			content[attr] = types.String{Value: value.String()}
		}
		state.Content.Elems = content
		state.Content.ElemType = types.StringType

		// Set id
		state.Id.Value = state.Dn.Value
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
