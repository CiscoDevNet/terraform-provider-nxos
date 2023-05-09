// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
)

// NxosProvider defines the provider implementation.
type NxosProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// NxosProviderData defines the data passed to resources and data sources.
type NxosProviderData struct {
	client  *nxos.Client
	devices map[string]string
}

// NxosProviderModel describes the provider data model.
type NxosProviderModel struct {
	Username types.String              `tfsdk:"username"`
	Password types.String              `tfsdk:"password"`
	URL      types.String              `tfsdk:"url"`
	Insecure types.Bool                `tfsdk:"insecure"`
	Retries  types.Int64               `tfsdk:"retries"`
	Devices  []NxosProviderModelDevice `tfsdk:"devices"`
}

type NxosProviderModelDevice struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

func (p *NxosProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nxos"
	resp.Version = p.version
}

func (p *NxosProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the NX-OS device account. This can also be set as the NXOS_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the NX-OS device account. This can also be set as the NXOS_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco NX-OS device. This can also be set as the NXOS_URL environment variable. If no URL is provided, the URL of the first device from the `devices` list is being used.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the NXOS_INSECURE environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the NXOS_RETRIES environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Device name.",
							Required:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: "URL of the Cisco NX-OS device.",
							Required:            true,
						},
					},
				},
			},
		},
	}
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
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("NXOS_USERNAME")
	} else {
		username = config.Username.ValueString()
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
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("NXOS_PASSWORD")
	} else {
		password = config.Password.ValueString()
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
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("NXOS_URL")
		if url == "" && len(config.Devices) > 0 {
			url = config.Devices[0].URL.ValueString()
		}
	} else {
		url = config.URL.ValueString()
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
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("NXOS_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("NXOS_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
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
		devices[device.Name.ValueString()] = device.URL.ValueString()
	}

	data := NxosProviderData{
		client:  &c,
		devices: devices,
	}

	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *NxosProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRestResource,
		NewBGPResource,
		NewBGPAddressFamilyResource,
		NewBGPGracefulRestartResource,
		NewBGPInstanceResource,
		NewBGPPeerResource,
		NewBGPPeerAddressFamilyResource,
		NewBGPPeerTemplateResource,
		NewBGPPeerTemplateAddressFamilyResource,
		NewBGPPeerTemplateMaxPrefixResource,
		NewBGPRouteControlResource,
		NewBGPVRFResource,
		NewBridgeDomainResource,
		NewDefaultQOSClassMapResource,
		NewDefaultQOSClassMapDSCPResource,
		NewDefaultQOSPolicyInterfaceInResource,
		NewDefaultQOSPolicyInterfaceInPolicyMapResource,
		NewDefaultQOSPolicyMapResource,
		NewDefaultQOSPolicyMapMatchClassMapResource,
		NewDefaultQOSPolicyMapMatchClassMapPoliceResource,
		NewDefaultQOSPolicyMapMatchClassMapSetQOSGroupResource,
		NewDHCPRelayAddressResource,
		NewDHCPRelayInterfaceResource,
		NewEthernetResource,
		NewEVPNResource,
		NewEVPNVNIResource,
		NewEVPNVNIRouteTargetResource,
		NewEVPNVNIRouteTargetDirectionResource,
		NewFeatureBFDResource,
		NewFeatureBGPResource,
		NewFeatureDHCPResource,
		NewFeatureEVPNResource,
		NewFeatureHMMResource,
		NewFeatureHSRPResource,
		NewFeatureInterfaceVLANResource,
		NewFeatureISISResource,
		NewFeatureLACPResource,
		NewFeatureLLDPResource,
		NewFeatureMACsecResource,
		NewFeatureNetflowResource,
		NewFeatureNVOverlayResource,
		NewFeatureOSPFResource,
		NewFeatureOSPFv3Resource,
		NewFeaturePIMResource,
		NewFeaturePTPResource,
		NewFeaturePVLANResource,
		NewFeatureSSHResource,
		NewFeatureTACACSResource,
		NewFeatureTelnetResource,
		NewFeatureUDLDResource,
		NewFeatureVNSegmentResource,
		NewFeatureVPCResource,
		NewHMMResource,
		NewHMMInstanceResource,
		NewHMMInterfaceResource,
		NewIPv4AccessListResource,
		NewIPv4AccessListEntryResource,
		NewIPv4AccessListPolicyEgressInterfaceResource,
		NewIPv4AccessListPolicyEgressInterfaceInstaceResource,
		NewIPv4AccessListPolicyIngressInterfaceResource,
		NewIPv4AccessListPolicyIngressInterfaceInstaceResource,
		NewIPv4InterfaceResource,
		NewIPv4InterfaceAddressResource,
		NewIPv4PrefixListRuleResource,
		NewIPv4PrefixListRuleEntryResource,
		NewIPv4VRFResource,
		NewISISResource,
		NewISISInstanceResource,
		NewISISInterfaceResource,
		NewISISVRFResource,
		NewLoopbackInterfaceResource,
		NewLoopbackInterfaceVRFResource,
		NewNTPServerResource,
		NewNVEInterfaceResource,
		NewNVEVNIResource,
		NewNVEVNIContainerResource,
		NewNVEVNIIngressReplicationResource,
		NewOSPFResource,
		NewOSPFAreaResource,
		NewOSPFAuthenticationResource,
		NewOSPFInstanceResource,
		NewOSPFInterfaceResource,
		NewOSPFVRFResource,
		NewPhysicalInterfaceResource,
		NewPhysicalInterfaceVRFResource,
		NewPIMResource,
		NewPIMAnycastRPResource,
		NewPIMAnycastRPPeerResource,
		NewPIMInstanceResource,
		NewPIMInterfaceResource,
		NewPIMSSMPolicyResource,
		NewPIMSSMRangeResource,
		NewPIMStaticRPResource,
		NewPIMStaticRPGroupListResource,
		NewPIMStaticRPPolicyResource,
		NewPIMVRFResource,
		NewQueuingQOSPolicyMapResource,
		NewQueuingQOSPolicyMapMatchClassMapResource,
		NewQueuingQOSPolicyMapMatchClassMapPriorityResource,
		NewQueuingQOSPolicyMapMatchClassMapRemainingBandwidthResource,
		NewQueuingQOSPolicySystemOutResource,
		NewQueuingQOSPolicySystemOutPolicyMapResource,
		NewSpanningTreeInterfaceResource,
		NewSubinterfaceResource,
		NewSubinterfaceVRFResource,
		NewSVIInterfaceResource,
		NewSVIInterfaceVRFResource,
		NewSystemResource,
		NewVRFResource,
		NewVRFAddressFamilyResource,
		NewVRFRouteTargetResource,
		NewVRFRouteTargetAddressFamilyResource,
		NewVRFRouteTargetDirectionResource,
		NewVRFRoutingResource,
	}
}

func (p *NxosProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRestDataSource,
		NewBGPDataSource,
		NewBGPAddressFamilyDataSource,
		NewBGPGracefulRestartDataSource,
		NewBGPInstanceDataSource,
		NewBGPPeerDataSource,
		NewBGPPeerAddressFamilyDataSource,
		NewBGPPeerTemplateDataSource,
		NewBGPPeerTemplateAddressFamilyDataSource,
		NewBGPPeerTemplateMaxPrefixDataSource,
		NewBGPRouteControlDataSource,
		NewBGPVRFDataSource,
		NewBridgeDomainDataSource,
		NewDefaultQOSClassMapDataSource,
		NewDefaultQOSClassMapDSCPDataSource,
		NewDefaultQOSPolicyInterfaceInDataSource,
		NewDefaultQOSPolicyInterfaceInPolicyMapDataSource,
		NewDefaultQOSPolicyMapDataSource,
		NewDefaultQOSPolicyMapMatchClassMapDataSource,
		NewDefaultQOSPolicyMapMatchClassMapPoliceDataSource,
		NewDefaultQOSPolicyMapMatchClassMapSetQOSGroupDataSource,
		NewDHCPRelayAddressDataSource,
		NewDHCPRelayInterfaceDataSource,
		NewEthernetDataSource,
		NewEVPNDataSource,
		NewEVPNVNIDataSource,
		NewEVPNVNIRouteTargetDataSource,
		NewEVPNVNIRouteTargetDirectionDataSource,
		NewFeatureBFDDataSource,
		NewFeatureBGPDataSource,
		NewFeatureDHCPDataSource,
		NewFeatureEVPNDataSource,
		NewFeatureHMMDataSource,
		NewFeatureHSRPDataSource,
		NewFeatureInterfaceVLANDataSource,
		NewFeatureISISDataSource,
		NewFeatureLACPDataSource,
		NewFeatureLLDPDataSource,
		NewFeatureMACsecDataSource,
		NewFeatureNetflowDataSource,
		NewFeatureNVOverlayDataSource,
		NewFeatureOSPFDataSource,
		NewFeatureOSPFv3DataSource,
		NewFeaturePIMDataSource,
		NewFeaturePTPDataSource,
		NewFeaturePVLANDataSource,
		NewFeatureSSHDataSource,
		NewFeatureTACACSDataSource,
		NewFeatureTelnetDataSource,
		NewFeatureUDLDDataSource,
		NewFeatureVNSegmentDataSource,
		NewFeatureVPCDataSource,
		NewHMMDataSource,
		NewHMMInstanceDataSource,
		NewHMMInterfaceDataSource,
		NewIPv4AccessListDataSource,
		NewIPv4AccessListEntryDataSource,
		NewIPv4AccessListPolicyEgressInterfaceDataSource,
		NewIPv4AccessListPolicyEgressInterfaceInstaceDataSource,
		NewIPv4AccessListPolicyIngressInterfaceDataSource,
		NewIPv4AccessListPolicyIngressInterfaceInstaceDataSource,
		NewIPv4InterfaceDataSource,
		NewIPv4InterfaceAddressDataSource,
		NewIPv4PrefixListRuleDataSource,
		NewIPv4PrefixListRuleEntryDataSource,
		NewIPv4VRFDataSource,
		NewISISDataSource,
		NewISISInstanceDataSource,
		NewISISInterfaceDataSource,
		NewISISVRFDataSource,
		NewLoopbackInterfaceDataSource,
		NewLoopbackInterfaceVRFDataSource,
		NewNTPServerDataSource,
		NewNVEInterfaceDataSource,
		NewNVEVNIDataSource,
		NewNVEVNIContainerDataSource,
		NewNVEVNIIngressReplicationDataSource,
		NewOSPFDataSource,
		NewOSPFAreaDataSource,
		NewOSPFAuthenticationDataSource,
		NewOSPFInstanceDataSource,
		NewOSPFInterfaceDataSource,
		NewOSPFVRFDataSource,
		NewPhysicalInterfaceDataSource,
		NewPhysicalInterfaceVRFDataSource,
		NewPIMDataSource,
		NewPIMAnycastRPDataSource,
		NewPIMAnycastRPPeerDataSource,
		NewPIMInstanceDataSource,
		NewPIMInterfaceDataSource,
		NewPIMSSMPolicyDataSource,
		NewPIMSSMRangeDataSource,
		NewPIMStaticRPDataSource,
		NewPIMStaticRPGroupListDataSource,
		NewPIMStaticRPPolicyDataSource,
		NewPIMVRFDataSource,
		NewQueuingQOSPolicyMapDataSource,
		NewQueuingQOSPolicyMapMatchClassMapDataSource,
		NewQueuingQOSPolicyMapMatchClassMapPriorityDataSource,
		NewQueuingQOSPolicyMapMatchClassMapRemainingBandwidthDataSource,
		NewQueuingQOSPolicySystemOutDataSource,
		NewQueuingQOSPolicySystemOutPolicyMapDataSource,
		NewSpanningTreeInterfaceDataSource,
		NewSubinterfaceDataSource,
		NewSubinterfaceVRFDataSource,
		NewSVIInterfaceDataSource,
		NewSVIInterfaceVRFDataSource,
		NewSystemDataSource,
		NewVRFDataSource,
		NewVRFAddressFamilyDataSource,
		NewVRFRouteTargetDataSource,
		NewVRFRouteTargetAddressFamilyDataSource,
		NewVRFRouteTargetDirectionDataSource,
		NewVRFRoutingDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NxosProvider{
			version: version,
		}
	}
}
