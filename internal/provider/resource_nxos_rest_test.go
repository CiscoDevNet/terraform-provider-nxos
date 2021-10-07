package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNxosRest_interface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosRestConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "id", "sys/intf/phys-[eth1/1]"),
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
				ResourceName:      "nxos_rest.l1PhysIf",
				ImportState:       true,
				ImportStateId:     "l1PhysIf:sys/intf/phys-[eth1/1]",
				ImportStateVerify: true,
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

func testAccCheckNxosRestDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(apiClient).Client

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "nxos_rest" {
			_, err := client.Get(rs.Primary.ID)
			if err == nil {
				return fmt.Errorf("Resource nxos_rest %s still exists", rs.Primary.ID)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckNxosRestStillExists(s *terraform.State) error {
	client := testAccProvider.Meta().(apiClient).Client

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "nxos_rest" {
			_, err := client.Get(rs.Primary.ID)
			if err != nil {
				return fmt.Errorf("Error retrieving resource nxos_rest %s", rs.Primary.ID)
			}

		} else {
			continue
		}
	}

	return nil
}
