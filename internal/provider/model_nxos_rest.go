package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
)

type Rest struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Delete    types.Bool   `tfsdk:"delete"`
	Content   types.Map    `tfsdk:"content"`
}

type RestDataSource struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Content   types.Map    `tfsdk:"content"`
}

func (data Rest) toBody(ctx context.Context) nxos.Body {
	body := nxos.Body{}

	className := data.ClassName.Value

	var content map[string]string
	data.Content.ElementsAs(ctx, &content, false)

	body.Str = fmt.Sprintf("{\"%s\":{\"attributes\":{}}}", className)
	for attr, value := range content {
		body = body.Set(fmt.Sprintf("%s.attributes.%s", className, attr), value)
	}

	return body
}
