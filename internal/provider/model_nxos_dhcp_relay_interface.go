// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type DHCPRelayInterface struct {
	Dn types.String `tfsdk:"id"`
	Id types.String `tfsdk:"interface_id"`
}

func (data DHCPRelayInterface) getDn() string {
	return fmt.Sprintf("sys/dhcp/inst/relayif-[%s]", data.Id.Value)
}

func (data DHCPRelayInterface) getClassName() string {
	return "dhcpRelayIf"
}

func (data DHCPRelayInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *DHCPRelayInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
}

func (data *DHCPRelayInterface) fromPlan(plan DHCPRelayInterface) {
}