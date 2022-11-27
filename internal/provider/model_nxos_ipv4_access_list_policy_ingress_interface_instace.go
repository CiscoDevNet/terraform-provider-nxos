// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type IPv4AccessListPolicyIngressInterfaceInstace struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Interface types.String `tfsdk:"interface_id"`
	Name      types.String `tfsdk:"name"`
}

func (data IPv4AccessListPolicyIngressInterfaceInstace) getDn() string {
	return fmt.Sprintf("sys/acl/ipv4/policy/ingress/intf-[%s]/acl", data.Interface.ValueString())
}

func (data IPv4AccessListPolicyIngressInterfaceInstace) getClassName() string {
	return "aclInst"
}

func (data IPv4AccessListPolicyIngressInterfaceInstace) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *IPv4AccessListPolicyIngressInterfaceInstace) fromBody(res gjson.Result) {
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
}

func (data *IPv4AccessListPolicyIngressInterfaceInstace) fromPlan(plan IPv4AccessListPolicyIngressInterfaceInstace) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Interface = plan.Interface
}
