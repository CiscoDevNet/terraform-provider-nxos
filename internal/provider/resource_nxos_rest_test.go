package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosRest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosRestConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "id", "sys/intf/phys-[eth1/1]"),
				),
			},
			{
				Config: testAccNxosRestConfig_child(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "class_name", "ipqosCMapInst"),
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "id", "sys/ipqos/dflt/c/name-[CM1]"),
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "children.0.content.val", "ef"),
				),
			},
			{
				Config: testAccNxosRestConfig_interface("Create description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "id", "sys/intf/phys-[eth1/1]"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "content.descr", "Create description"),
				),
			},
			{
				ResourceName:  "nxos_rest.l1PhysIf",
				ImportState:   true,
				ImportStateId: "l1PhysIf:sys/intf/phys-[eth1/1]",
			},
			{
				Config: testAccNxosRestConfig_interface("Updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "content.descr", "Updated description"),
				),
			},
		},
	})
}

func testAccNxosRestConfig_empty() string {
	return `
	resource "nxos_rest" "l1PhysIf" {
		dn = "sys/intf/phys-[eth1/1]"
		class_name = "l1PhysIf"
	}
	`
}

func testAccNxosRestConfig_interface(description string) string {
	return fmt.Sprintf(`
	resource "nxos_rest" "l1PhysIf" {
		dn = "sys/intf/phys-[eth1/1]"
		class_name = "l1PhysIf"
		content = {
			id = "eth1/1"
			descr = "%[1]s"
		}
	}
	`, description)
}

func testAccNxosRestConfig_child() string {
	return `
	resource "nxos_rest" "ipqosCMapInst" {
		dn = "sys/ipqos/dflt/c/name-[CM1]"
		class_name = "ipqosCMapInst"
		content = {
			name = "CM1"
		}
		children = [
            {
                rn = "dscp-ef"
                class_name = "ipqosDscp"
				content = {
					val = "ef"
				}
			}
		]
	}
	`
}
