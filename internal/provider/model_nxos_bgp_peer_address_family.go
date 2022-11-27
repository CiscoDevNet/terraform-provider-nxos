// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeerAddressFamily struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	Asn        types.String `tfsdk:"asn"`
	Vrf_name   types.String `tfsdk:"vrf"`
	Addr       types.String `tfsdk:"address"`
	Type       types.String `tfsdk:"address_family"`
	Ctrl       types.String `tfsdk:"control"`
	SendComExt types.String `tfsdk:"send_community_extended"`
	SendComStd types.String `tfsdk:"send_community_standard"`
}

func (data BGPPeerAddressFamily) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]/af-[%s]", data.Vrf_name.ValueString(), data.Addr.ValueString(), data.Type.ValueString())
}

func (data BGPPeerAddressFamily) getClassName() string {
	return "bgpPeerAf"
}

func (data BGPPeerAddressFamily) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.ValueString()).
		Set("ctrl", data.Ctrl.ValueString()).
		Set("sendComExt", data.SendComExt.ValueString()).
		Set("sendComStd", data.SendComStd.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeerAddressFamily) fromBody(res gjson.Result) {
	data.Type = types.StringValue(res.Get("*.attributes.type").String())
	data.Ctrl = types.StringValue(res.Get("*.attributes.ctrl").String())
	data.SendComExt = types.StringValue(res.Get("*.attributes.sendComExt").String())
	data.SendComStd = types.StringValue(res.Get("*.attributes.sendComStd").String())
}

func (data *BGPPeerAddressFamily) fromPlan(plan BGPPeerAddressFamily) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Asn = plan.Asn
	data.Vrf_name = plan.Vrf_name
	data.Addr = plan.Addr
}
