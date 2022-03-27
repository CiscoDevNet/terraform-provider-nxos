// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map.test", "name", "PM1"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map.test", "match_type", "match-any"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapConfig = `

resource "nxos_default_qos_policy_map" "test" {
  name = "PM1"
  match_type = "match-any"
}

data "nxos_default_qos_policy_map" "test" {
  name = "PM1"
  depends_on = [nxos_default_qos_policy_map.test]
}
`