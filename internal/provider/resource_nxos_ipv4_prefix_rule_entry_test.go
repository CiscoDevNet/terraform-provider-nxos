// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosIPv4PrefixRuleEntry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv4PrefixRuleEntryConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "rule_name", "RULE1"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "order", "10"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "criteria", "exact"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "prefix", "192.168.1.0/24"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "from_range", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_prefix_rule_entry.test", "to_range", "24"),
				),
			},
			{
				ResourceName:  "nxos_ipv4_prefix_rule_entry.test",
				ImportState:   true,
				ImportStateId: "sys/rpm/pfxlistv4-[RULE1]/ent-[10]",
			},
		},
	})
}

func testAccNxosIPv4PrefixRuleEntryConfig_minimum() string {
	return `
	resource "nxos_ipv4_prefix_rule_entry" "test" {
		rule_name = "RULE1"
		order = 10
	}
	`
}

func testAccNxosIPv4PrefixRuleEntryConfig_all() string {
	return `
	resource "nxos_ipv4_prefix_rule_entry" "test" {
		rule_name = "RULE1"
		action = "permit"
		order = 10
		criteria = "exact"
		prefix = "192.168.1.0/24"
		from_range = 0
		to_range = 24
	}
	`
}
