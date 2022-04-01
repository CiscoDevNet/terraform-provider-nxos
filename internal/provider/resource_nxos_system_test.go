// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosSystem(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosSystemConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_system.test", "name", "LEAF1"),
				),
			},
			{
				ResourceName:  "nxos_system.test",
				ImportState:   true,
				ImportStateId: "sys",
			},
		},
	})
}

func testAccNxosSystemConfig_minimum() string {
	return `
	resource "nxos_system" "test" {
	}
	`
}

func testAccNxosSystemConfig_all() string {
	return `
	resource "nxos_system" "test" {
		name = "LEAF1"
	}
	`
}