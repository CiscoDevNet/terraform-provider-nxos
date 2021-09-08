package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netascode/go-nxos"
)

func resourceNxosRest() *schema.Resource {
	return &schema.Resource{
		Description: "Manages NXOS Model Objects via REST API calls. This resource can only manage a single API object. It is able to read the state and therefore reconcile configuration drift.",

		CreateContext: resourceNxosRestCreate,
		UpdateContext: resourceNxosRestUpdate,
		ReadContext:   resourceNxosRestRead,
		DeleteContext: resourceNxosRestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceNxosRestImport,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The distinguished name of the object.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dn": {
				Type:        schema.TypeString,
				Description: "Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].",
				Required:    true,
				ForceNew:    true,
			},
			"class_name": {
				Type:        schema.TypeString,
				Description: "Which class object is being created. (Make sure there is no colon in the classname)",
				Required:    true,
				ForceNew:    true,
			},
			"content": {
				Type:        schema.TypeMap,
				Description: "Map of key-value pairs that need to be passed to the Model object as parameters.",
				Optional:    true,
				Computed:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					content := d.Get("content")
					contentStrMap := toStrMap(content.(map[string]interface{}))
					key := strings.Split(k, ".")[1]
					if _, ok := contentStrMap[key]; ok {
						return false
					}
					return true
				},
			},
		},
	}
}

func resourceNxosRestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	client := meta.(apiClient).Client

	for attempts := 0; ; attempts++ {
		res, err := client.GetDn(d.Get("dn").(string), nxos.Query("rsp-prop-include", "config-only"))

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
			return nil
		}

		className := d.Get("class_name").(string)
		dn := d.Get("dn").(string)
		d.SetId(dn)

		newContent := make(map[string]interface{})

		for attr, value := range res.Get(className + ".attributes").Map() {
			newContent[attr] = value.String()
		}

		d.Set("content", newContent)
		break
	}

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceNxosRestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Create", d.Id())
	client := meta.(apiClient).Client

	for attempts := 0; ; attempts++ {
		body, _ := prepareBody(d)
		_, err := client.Post(d.Get("dn").(string), body.Str)

		if err == nil {
			break
		}
		if ok := backoff(attempts, meta.(apiClient).Retries); !ok {
			return diag.FromErr(err)
		}
		log.Printf("[ERROR] Failed to update object: %s, retries: %v", err, attempts)
	}

	log.Printf("[DEBUG] %s: Create finished successfully", d.Id())
	return resourceNxosRestRead(ctx, d, meta)
}

func resourceNxosRestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Update", d.Id())
	client := meta.(apiClient).Client

	for attempts := 0; ; attempts++ {
		body, _ := prepareBody(d)
		_, err := client.Post(d.Get("dn").(string), body.Str)

		if err == nil {
			break
		}
		if ok := backoff(attempts, meta.(apiClient).Retries); !ok {
			return diag.FromErr(err)
		}
		log.Printf("[ERROR] Failed to update object: %s, retries: %v", err, attempts)
	}

	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceNxosRestRead(ctx, d, meta)
}

func resourceNxosRestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	client := meta.(apiClient).Client

	for attempts := 0; ; attempts++ {
		res, err := client.DeleteDn(d.Get("dn").(string))

		if err == nil {
			break
		}
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode == "107" {
			break
		}
		if ok := backoff(attempts, meta.(apiClient).Retries); !ok {
			return diag.FromErr(err)
		}
		log.Printf("[ERROR] Failed to delete object: %s, retries: %v", err, attempts)
	}

	d.SetId("")
	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	return nil
}

func resourceNxosRestImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())

	parts := strings.SplitN(d.Id(), ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%s), expected class_name:dn", d.Id())
	}

	d.Set("dn", parts[1])
	d.Set("class_name", parts[0])
	d.SetId(parts[1])

	if diags := resourceNxosRestRead(ctx, d, meta); diags.HasError() {
		return nil, fmt.Errorf("Could not read object when importing: %s", diags[0].Summary)
	}

	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{d}, nil
}
