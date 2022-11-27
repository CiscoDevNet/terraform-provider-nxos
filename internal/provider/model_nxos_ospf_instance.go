// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type OSPFInstance struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
	Name    types.String `tfsdk:"name"`
}

func (data OSPFInstance) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]", data.Name.ValueString())
}

func (data OSPFInstance) getClassName() string {
	return "ospfInst"
}

func (data OSPFInstance) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.ValueString()).
		Set("name", data.Name.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *OSPFInstance) fromBody(res gjson.Result) {
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
}

func (data *OSPFInstance) fromPlan(plan OSPFInstance) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
