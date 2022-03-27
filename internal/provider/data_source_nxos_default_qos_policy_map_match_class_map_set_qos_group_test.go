// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupPrerequisitesConfig + testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_set_qos_group.test", "qos_group_id", "1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipqos/dflt/p/name-[PM1]"
  class_name = "ipqosPMapInst"
  content = {
      name = "PM1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipqos/dflt/c/name-[Voice]"
  class_name = "ipqosCMapInst"
  content = {
      name = "Voice"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ipqos/dflt/p/name-[PM1]/cmap-[Voice]"
  class_name = "ipqosMatchCMap"
  content = {
      name = "Voice"
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapSetQOSGroupConfig = `

resource "nxos_default_qos_policy_map_match_class_map_set_qos_group" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
  qos_group_id = 1
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_default_qos_policy_map_match_class_map_set_qos_group" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
  depends_on = [nxos_default_qos_policy_map_match_class_map_set_qos_group.test]
}
`