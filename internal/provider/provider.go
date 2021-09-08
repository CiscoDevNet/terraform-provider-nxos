package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netascode/go-nxos"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("NXOS_USERNAME", nil),
					Description: "Username for the NXOS device account. This can also be set as the NXOS_USERNAME environment variable.",
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("NXOS_PASSWORD", nil),
					Description: "Password for the NXOS device account. This can also be set as the NXOS_PASSWORD environment variable.",
				},
				"url": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("NXOS_URL", nil),
					Description: "URL of the Cisco NXOS device. This can also be set as the NXOS_URL environment variable.",
				},
				"insecure": {
					Type:     schema.TypeBool,
					Optional: true,
					DefaultFunc: func() (interface{}, error) {
						if v := os.Getenv("NXOS_INSECURE"); v != "" {
							return strconv.ParseBool(v)
						}
						return true, nil
					},
					Description: "Allow insecure HTTPS client. This can also be set as the NXOS_INSECURE environment variable. Defaults to `true`.",
				},
				"retries": {
					Type:     schema.TypeInt,
					Optional: true,
					DefaultFunc: func() (interface{}, error) {
						if v := os.Getenv("NXOS_RETRIES"); v != "" {
							return strconv.Atoi(v)
						}
						return 3, nil
					},
					ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
						v := val.(int)
						if v < 0 || v > 9 {
							errs = append(errs, fmt.Errorf("%q must be between 0 and 9 inclusive, got: %d", key, v))
						}
						return
					},
					Description: "Number of retries for REST API calls. This can also be set as the NXOS_RETRIES environment variable. Defaults to `3`.",
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"nxos_rest": dataSourceNxosRest(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"nxos_rest": resourceNxosRest(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	Username   string
	Password   string
	URL        string
	IsInsecure bool
	Retries    int
	Client     nxos.Client
}

func (c apiClient) Valid() diag.Diagnostics {

	if c.Username == "" {
		return diag.FromErr(fmt.Errorf("Username must be provided for the NXOS provider"))
	}

	if c.Password == "" {
		return diag.FromErr(fmt.Errorf("Password must be provided for the NXOS provider"))
	}

	if c.URL == "" {
		return diag.FromErr(fmt.Errorf("The URL must be provided for the NXOS provider"))
	}

	return nil
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		cl := apiClient{
			Username:   d.Get("username").(string),
			Password:   d.Get("password").(string),
			URL:        d.Get("url").(string),
			IsInsecure: d.Get("insecure").(bool),
			Retries:    d.Get("retries").(int),
		}

		if diag := cl.Valid(); diag != nil {
			return nil, diag
		}

		cl.Client, _ = nxos.NewClient(cl.URL, cl.Username, cl.Password, cl.IsInsecure)

		return cl, nil
	}
}
