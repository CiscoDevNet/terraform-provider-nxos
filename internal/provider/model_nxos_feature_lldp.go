// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureLLDP struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureLLDP) getDn() string {
	return "sys/fm/lldp"
}

func (data FeatureLLDP) getClassName() string {
	return "fmLldp"
}

func (data FeatureLLDP) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureLLDP) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
}

func (data *FeatureLLDP) fromPlan(plan FeatureLLDP) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
