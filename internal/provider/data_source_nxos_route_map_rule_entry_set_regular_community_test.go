// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosRouteMapRuleEntrySetRegularCommunity(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosRouteMapRuleEntrySetRegularCommunityPrerequisitesConfig + testAccDataSourceNxosRouteMapRuleEntrySetRegularCommunityConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_route_map_rule_entry_set_regular_community.test", "additive", "disabled"),
					resource.TestCheckResourceAttr("data.nxos_route_map_rule_entry_set_regular_community.test", "no_community", "disabled"),
					resource.TestCheckResourceAttr("data.nxos_route_map_rule_entry_set_regular_community.test", "set_criteria", "none"),
				),
			},
		},
	})
}

const testAccDataSourceNxosRouteMapRuleEntrySetRegularCommunityPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/rpm/rtmap-[RULE1]"
  class_name = "rtmapRule"
  content = {
      name = "RULE1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/rpm/rtmap-[RULE1]/ent-[0]"
  class_name = "rtmapEntry"
  content = {
      name = "RULE1"
      order = "0"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

`

const testAccDataSourceNxosRouteMapRuleEntrySetRegularCommunityConfig = `

resource "nxos_route_map_rule_entry_set_regular_community" "test" {
  rule_name = "RULE1"
  order = 10
  additive = "disabled"
  no_community = "disabled"
  set_criteria = "none"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_route_map_rule_entry_set_regular_community" "test" {
  rule_name = "RULE1"
  order = 10
  depends_on = [nxos_route_map_rule_entry_set_regular_community.test]
}
`
