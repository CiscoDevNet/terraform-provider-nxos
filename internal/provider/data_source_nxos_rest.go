package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
)

type dataSourceRestType struct{}

func (t dataSourceRestType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read one NX-OS object.",

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

func (t dataSourceRestType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceRest{
		provider: provider,
	}, diags
}

type dataSourceRest struct {
	provider provider
}

func (d dataSourceRest) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state Rest

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.Value))

	res, err := d.provider.client.GetDn(config.Dn.Value, nxos.OverrideUrl(d.provider.devices[config.Device.Value]))
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
