// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPPeerAddressFamilyPrefixListControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPeerAddressFamilyPrefixListControlPrerequisitesConfig + testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "address", "192.168.0.1"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "direction", "in"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family_prefix_list_control.test", "list", "PREFIX_LIST1"),
				),
			},
			{
				ResourceName:  "nxos_bgp_peer_address_family_prefix_list_control.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/pfxctrl-[in]",
			},
		},
	})
}

const testAccNxosBGPPeerAddressFamilyPrefixListControlPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]"
  class_name = "bgpPeer"
  content = {
      addr = "192.168.0.1"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]"
  class_name = "bgpPeerAf"
  content = {
      addr = "192.168.0.1"
      asn = "65001"
      address_family = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

func testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_minimum() string {
	return `
	resource "nxos_bgp_peer_address_family_prefix_list_control" "test" {
		asn = "65001"
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
		direction = "in"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

func testAccNxosBGPPeerAddressFamilyPrefixListControlConfig_all() string {
	return `
	resource "nxos_bgp_peer_address_family_prefix_list_control" "test" {
		asn = "65001"
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
		direction = "in"
		list = "PREFIX_LIST1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}