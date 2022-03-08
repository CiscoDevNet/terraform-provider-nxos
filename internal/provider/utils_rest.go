package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/netascode/go-nxos"
)

func prepareBody(ctx context.Context, d Rest) (nxos.Body, diag.Diagnostics) {
	body := nxos.Body{}

	className := d.ClassName.Value

	body.Str = fmt.Sprintf("{\"%s\":{\"attributes\":{}}}", className)
	var content map[string]string
	d.Content.ElementsAs(ctx, &content, false)
	for attr, value := range content {
		body = body.Set(fmt.Sprintf("%s.attributes.%s", className, attr), value)
	}
	return body, nil
}
