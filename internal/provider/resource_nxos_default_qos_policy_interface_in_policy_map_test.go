// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosDefaultQOSPolicyInterfaceInPolicyMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosDefaultQOSPolicyInterfaceInPolicyMapPrerequisitesConfig + testAccNxosDefaultQOSPolicyInterfaceInPolicyMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_default_qos_policy_interface_in_policy_map.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_default_qos_policy_interface_in_policy_map.test", "policy_map_name", "PM1"),
				),
			},
			{
				ResourceName:  "nxos_default_qos_policy_interface_in_policy_map.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/dflt/policy/in/intf-[eth1/10]/pmap",
			},
		},
	})
}

const testAccNxosDefaultQOSPolicyInterfaceInPolicyMapPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipqos/dflt/p/name-[PM1]"
  class_name = "ipqosPMapInst"
  content = {
      name = "PM1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipqos/dflt/policy/in/intf-[eth1/10]"
  class_name = "ipqosIf"
  content = {
      name = "eth1/10"
  }
}

`

func testAccNxosDefaultQOSPolicyInterfaceInPolicyMapConfig_minimum() string {
	return `
	resource "nxos_default_qos_policy_interface_in_policy_map" "test" {
		interface_id = "eth1/10"
		policy_map_name = "PM1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosDefaultQOSPolicyInterfaceInPolicyMapConfig_all() string {
	return `
	resource "nxos_default_qos_policy_interface_in_policy_map" "test" {
		interface_id = "eth1/10"
		policy_map_name = "PM1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}