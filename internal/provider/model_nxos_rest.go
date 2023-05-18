package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type RestModel struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Delete    types.Bool   `tfsdk:"delete"`
	Content   types.Map    `tfsdk:"content"`
}

type RestDataSourceModel struct {
	Device    types.String `tfsdk:"device"`
	Id        types.String `tfsdk:"id"`
	Dn        types.String `tfsdk:"dn"`
	ClassName types.String `tfsdk:"class_name"`
	Content   types.Map    `tfsdk:"content"`
}

func (data RestModel) toBody(ctx context.Context) nxos.Body {
	body := nxos.Body{}

	className := data.ClassName.ValueString()

	var content map[string]string
	data.Content.ElementsAs(ctx, &content, false)

	body.Str = fmt.Sprintf("{\"%s\":{\"attributes\":{}}}", className)
	for attr, value := range content {
		body = body.Set(fmt.Sprintf("%s.attributes.%s", className, attr), value)
	}

	return body
}

func (data *RestModel) fromBody(ctx context.Context, res gjson.Result) {
	if !res.Exists() {
		data.Id = types.StringNull()
		return
	}

	stateContent := data.Content
	newContent := make(map[string]attr.Value)
	for attr, value := range res.Get(data.ClassName.ValueString() + ".attributes").Map() {
		if _, ok := stateContent.Elements()[attr]; ok {
			newContent[attr] = types.StringValue(value.Str)
		}
	}
	if len(newContent) > 0 {
		data.Content = types.MapValueMust(data.Content.ElementType(ctx), newContent)
	} else {
		data.Content = types.MapNull(data.Content.ElementType(ctx))
	}
}
