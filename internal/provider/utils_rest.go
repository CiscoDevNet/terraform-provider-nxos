package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/netascode/go-nxos"
)

func prepareBody(d *schema.ResourceData) (nxos.Body, error) {
	body := nxos.Body{}

	className := d.Get("class_name").(string)
	content := d.Get("content")
	contentStrMap := toStrMap(content.(map[string]interface{}))

	for attr, value := range contentStrMap {
		body = body.Set(fmt.Sprintf("%s.attributes.%s", className, attr), value)
	}
	return body, nil
}
