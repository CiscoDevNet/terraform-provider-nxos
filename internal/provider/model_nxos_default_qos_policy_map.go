// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type DefaultQOSPolicyMap struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	MatchType types.String `tfsdk:"match_type"`
}

func (data DefaultQOSPolicyMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/p/name-[%s]", data.Name.ValueString())
}

func (data DefaultQOSPolicyMap) getClassName() string {
	return "ipqosPMapInst"
}

func (data DefaultQOSPolicyMap) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.ValueString()).
		Set("matchType", data.MatchType.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *DefaultQOSPolicyMap) fromBody(res gjson.Result) {
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
	data.MatchType = types.StringValue(res.Get("*.attributes.matchType").String())
}

func (data *DefaultQOSPolicyMap) fromPlan(plan DefaultQOSPolicyMap) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}
