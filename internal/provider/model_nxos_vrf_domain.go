// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFDomain struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Vrf    types.String `tfsdk:"vrf"`
	Rd     types.String `tfsdk:"rd"`
}

// Helper function to get Dn, should be moved to helpers
func Replacer(format string, args ...string) string {
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}

func (data VRFDomain) getDn() string {
	return Replacer("sys/inst-[|vrf|]/dom-[|vrf|]", "|vrf|", data.Vrf.Value)
}

func (data VRFDomain) getClassName() string {
	return "rtctrlDom"
}

func (data VRFDomain) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("rd", data.Rd.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFDomain) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Rd.Value = res.Get("*.attributes.rd").String()
}

func (data *VRFDomain) fromPlan(plan VRFDomain) {
	data.Device = plan.Device
	data.Vrf.Value = plan.Vrf.Value
}
