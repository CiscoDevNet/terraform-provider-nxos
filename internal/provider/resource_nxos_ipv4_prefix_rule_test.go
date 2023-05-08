// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosIPv4PrefixRule(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv4PrefixRuleConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule.test", "name", "RULE1"),
				),
			},
			{
				ResourceName:  "nxos_ipv4_prefix_rule.test",
				ImportState:   true,
				ImportStateId: "sys/rpm/pfxlistv4-[RULE1]",
			},
		},
	})
}

func testAccNxosIPv4PrefixRuleConfig_minimum() string {
	return `
	resource "nxos_ipv4_prefix_rule" "test" {
		name = "RULE1"
	}
	`
}

func testAccNxosIPv4PrefixRuleConfig_all() string {
	return `
	resource "nxos_ipv4_prefix_rule" "test" {
		name = "RULE1"
	}
	`
}
