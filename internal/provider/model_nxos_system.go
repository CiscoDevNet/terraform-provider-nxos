// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type System struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
}

func (data System) getDn() string {
	return "sys"
}

func (data System) getClassName() string {
	return "topSystem"
}

func (data System) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *System) fromBody(res gjson.Result) {
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
}

func (data *System) fromPlan(plan System) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
