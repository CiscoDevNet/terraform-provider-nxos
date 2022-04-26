// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosIPv4Interface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosIPv4InterfacePrerequisitesConfig + testAccDataSourceNxosIPv4InterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface.test", "drop_glean", "disabled"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface.test", "forward", "disabled"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface.test", "unnumbered", "unspecified"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface.test", "urpf", "disabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4InterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipv4/inst/dom-[default]"
  class_name = "ipv4Dom"
  content = {
  }
}

`

const testAccDataSourceNxosIPv4InterfaceConfig = `

resource "nxos_ipv4_interface" "test" {
  vrf = "default"
  interface_id = "eth1/10"
  drop_glean = "disabled"
  forward = "disabled"
  unnumbered = "unspecified"
  urpf = "disabled"
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_ipv4_interface" "test" {
  vrf = "default"
  interface_id = "eth1/10"
  depends_on = [nxos_ipv4_interface.test]
}
`
