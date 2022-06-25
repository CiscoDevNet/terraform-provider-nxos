// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type IPv4AccessListPolicyEgressInterfaceInstace struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Interface types.String `tfsdk:"interface_id"`
	Name      types.String `tfsdk:"name"`
}

func (data IPv4AccessListPolicyEgressInterfaceInstace) getDn() string {
	return fmt.Sprintf("sys/acl/ipv4/policy/egress/intf-[%s]/acl", data.Interface.Value)
}

func (data IPv4AccessListPolicyEgressInterfaceInstace) getClassName() string {
	return "aclInst"
}

func (data IPv4AccessListPolicyEgressInterfaceInstace) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *IPv4AccessListPolicyEgressInterfaceInstace) fromBody(res gjson.Result) {
	data.Name.Value = res.Get("*.attributes.name").String()
}

func (data *IPv4AccessListPolicyEgressInterfaceInstace) fromPlan(plan IPv4AccessListPolicyEgressInterfaceInstace) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Interface.Value = plan.Interface.Value
}
