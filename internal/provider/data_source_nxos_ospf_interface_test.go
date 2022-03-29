// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosOSPFInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosOSPFInterfacePrerequisitesConfig + testAccDataSourceNxosOSPFInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "advertise_secondaries", "false"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "area", "0.0.0.10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "bfd", "disabled"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "cost", "1000"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "dead_interval", "60"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "hello_interval", "15"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "network_type", "p2p"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "passive", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_ospf_interface.test", "priority", "10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosOSPFInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bfd"
  class_name = "fmBfd"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[default]"
  class_name = "ospfDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosOSPFInterfaceConfig = `

resource "nxos_ospf_interface" "test" {
  instance_name = "OSPF1"
  vrf_name = "default"
  interface_id = "eth1/10"
  advertise_secondaries = false
  area = "0.0.0.10"
  bfd = "disabled"
  cost = 1000
  dead_interval = 60
  hello_interval = 15
  network_type = "p2p"
  passive = "enabled"
  priority = 10
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_ospf_interface" "test" {
  instance_name = "OSPF1"
  vrf_name = "default"
  interface_id = "eth1/10"
  depends_on = [nxos_ospf_interface.test]
}
`
