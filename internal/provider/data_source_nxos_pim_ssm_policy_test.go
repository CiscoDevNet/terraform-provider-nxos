// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMSSMPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPIMSSMPolicyPrerequisitesConfig + testAccDataSourceNxosPIMSSMPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_ssm_policy.test", "name", "SSM"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMSSMPolicyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/pim"
  class_name = "pimEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosPIMSSMPolicyConfig = `

resource "nxos_pim_ssm_policy" "test" {
  vrf_name = "default"
  name = "SSM"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_pim_ssm_policy" "test" {
  vrf_name = "default"
  depends_on = [nxos_pim_ssm_policy.test]
}
`