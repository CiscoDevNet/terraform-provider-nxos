package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosRest_interface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosRestConfigInterface,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_rest.int", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("data.nxos_rest.int", "content.id", "eth1/1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosRestConfigInterface = `
data "nxos_rest" "int" {
  dn = "sys/intf/phys-[eth1/1]"
}
`
