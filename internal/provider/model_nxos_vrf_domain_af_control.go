// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFDomainAfControl struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Vrf     types.String `tfsdk:"vrf"`
	Af_type types.String `tfsdk:"af"`
	Type    types.String `tfsdk:"rt_type"`
}

func (data VRFDomainAfControl) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]", data.Vrf.Value, data.Af_type.Value, data.Type.Value)
}

func (data VRFDomainAfControl) getClassName() string {
	return "rtctrlAfCtrl"
}

func (data VRFDomainAfControl) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFDomainAfControl) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Type.Value = res.Get("*.attributes.type").String()
}

func (data *VRFDomainAfControl) fromPlan(plan VRFDomainAfControl) {
	data.Device = plan.Device
	data.Vrf.Value = plan.Vrf.Value
	data.Af_type.Value = plan.Af_type.Value
}
