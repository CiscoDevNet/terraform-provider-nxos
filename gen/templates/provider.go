//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure NxosProvider satisfies various provider interfaces.
var _ provider.Provider = &NxosProvider{}
var _ provider.ProviderWithMetadata = &NxosProvider{}

// NxosProvider defines the provider implementation.
type NxosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NxosProviderData defines the data passed to resources and data sources.
type NxosProviderData struct {
	client *nxos.Client
	devices map[string]string
}

// NxosProviderModel  describes the provider data model.
type NxosProviderModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	URL      types.String `tfsdk:"url"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Retries  types.Int64  `tfsdk:"retries"`
	Devices  []NxosProviderModelDevice  `tfsdk:"devices"`
}

type NxosProviderModelDevice struct {
	Name types.String `tfsdk:"name"`
	URL      types.String `tfsdk:"url"`
}

func (p *NxosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nxos"
	resp.Version = p.version
}

func (p *NxosProvider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				MarkdownDescription: "Username for the NX-OS device account. This can also be set as the NXOS_USERNAME environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"password": {
				MarkdownDescription: "Password for the NX-OS device account. This can also be set as the NXOS_PASSWORD environment variable.",
				Type:                types.StringType,
				Optional:            true,
				Sensitive:           true,
			},
			"url": {
				MarkdownDescription: "URL of the Cisco NX-OS device. This can also be set as the NXOS_URL environment variable. If no URL is provided, the URL of the first device from the `devices` list is being used.",
				Type:                types.StringType,
				Optional:            true,
			},
			"insecure": {
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the NXOS_INSECURE environment variable. Defaults to `true`.",
				Type:                types.BoolType,
				Optional:            true,
			},
			"retries": {
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the NXOS_RETRIES environment variable. Defaults to `3`.",
				Type:                types.Int64Type,
				Optional:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 9),
				},
			},
			"devices": {
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "Device name.",
						Type:                types.StringType,
						Required:            true,
					},
					"url": {
						MarkdownDescription: "URL of the Cisco NX-OS device.",
						Type:                types.StringType,
						Required:            true,
					},
				}),
			},
		},
	}, nil
}

func (p *NxosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config NxosProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Null {
		username = os.Getenv("NXOS_USERNAME")
	} else {
		username = config.Username.Value
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Null {
		password = os.Getenv("NXOS_PASSWORD")
	} else {
		password = config.Password.Value
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.Null {
		url = os.Getenv("NXOS_URL")
		if url == "" && len(config.Devices) > 0 {
			url = config.Devices[0].URL.Value
		}
	} else {
		url = config.URL.Value
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.Null {
		insecureStr := os.Getenv("NXOS_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.Value
	}

	var retries int64
	if config.Retries.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.Null {
		retriesStr := os.Getenv("NXOS_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.Value
	}

	// Create a new NX-OS client and set it to the provider client
	c, err := nxos.NewClient(url, username, password, insecure, nxos.MaxRetries(int(retries)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create nxos client:\n\n"+err.Error(),
		)
		return
	}

	devices := make(map[string]string)
    for _, device := range config.Devices {
		devices[device.Name.Value] = device.URL.Value
	}

	data := NxosProviderData{
		client: &c,
		devices: devices,
	}

	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *NxosProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRestResource,
		{{- range .}}
		New{{camelCase .Name}}Resource,
		{{- end}}
	}
}

func (p *NxosProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRestDataSource,
		{{- range .}}
		New{{camelCase .Name}}DataSource,
		{{- end}}
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NxosProvider{
			version: version,
		}
	}
}
