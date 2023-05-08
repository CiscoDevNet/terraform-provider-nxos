// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeerAddressFamilyPrefixListControl struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Asn       types.String `tfsdk:"asn"`
	Vrf_name  types.String `tfsdk:"vrf"`
	Addr      types.String `tfsdk:"address"`
	Type      types.String `tfsdk:"address_family"`
	Direction types.String `tfsdk:"direction"`
	List      types.String `tfsdk:"list"`
}

func (data BGPPeerAddressFamilyPrefixListControl) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]/af-[%s]/pfxctrl-[%s]", data.Vrf_name.ValueString(), data.Addr.ValueString(), data.Type.ValueString(), data.Direction.ValueString())
}

func (data BGPPeerAddressFamilyPrefixListControl) getClassName() string {
	return "bgpPfxCtrlP"
}

func (data BGPPeerAddressFamilyPrefixListControl) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("direction", data.Direction.ValueString()).
		Set("list", data.List.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeerAddressFamilyPrefixListControl) fromBody(res gjson.Result) {
	data.Direction = types.StringValue(res.Get("*.attributes.direction").String())
	data.List = types.StringValue(res.Get("*.attributes.list").String())
}

func (data *BGPPeerAddressFamilyPrefixListControl) fromPlan(plan BGPPeerAddressFamilyPrefixListControl) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Asn = plan.Asn
	data.Vrf_name = plan.Vrf_name
	data.Addr = plan.Addr
	data.Type = plan.Type
}