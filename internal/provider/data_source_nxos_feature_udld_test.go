// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosFeatureUDLD(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosFeatureUDLDConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_feature_udld.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosFeatureUDLDConfig = `

resource "nxos_feature_udld" "test" {
  admin_state = "enabled"
}

data "nxos_feature_udld" "test" {
  depends_on = [nxos_feature_udld.test]
}
`