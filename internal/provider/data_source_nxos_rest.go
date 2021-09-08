package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNxosRest() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can read one NXOS object.",

		ReadContext: dataSourceNxosRestRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The distinguished name of the object.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dn": {
				Type:        schema.TypeString,
				Description: "Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].",
				Required:    true,
			},
			"class_name": {
				Type:        schema.TypeString,
				Description: "Class name of object being retrieved.",
				Computed:    true,
			},
			"content": {
				Type:        schema.TypeMap,
				Description: "Map of key-value pairs which represents the attributes of object being retrieved.",
				Computed:    true,
			},
		},
	}
}

func dataSourceNxosRestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	client := meta.(apiClient).Client

	for attempts := 0; ; attempts++ {
		res, err := client.GetDn(d.Get("dn").(string))

		if err != nil {
			if ok := backoff(attempts, meta.(apiClient).Retries); !ok {
				return diag.FromErr(err)
			}
			log.Printf("[ERROR] Failed to read object: %s, retries: %v", err, attempts)
			continue
		}

		// Check if we received an empty response without errors -> object has been deleted
		if !res.Exists() && err == nil {
			d.SetId("")
			break
		}

		// Set class_name
		var className string
		for class := range res.Map() {
			className = class
		}
		d.Set("class_name", className)

		// Set content
		content := make(map[string]interface{})

		for attr, value := range res.Get(className + ".attributes").Map() {
			content[attr] = value.String()
		}
		d.Set("content", content)

		// Set id
		d.SetId(d.Get("dn").(string))
		break
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}
