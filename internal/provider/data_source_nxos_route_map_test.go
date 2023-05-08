// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosRouteMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosRouteMapConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceNxosRouteMapConfig = `

resource "nxos_route_map" "test" {
}

data "nxos_route_map" "test" {
  depends_on = [nxos_route_map.test]
}
`
