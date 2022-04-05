// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosVRFDomainAfControl(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVRFDomainAfControlConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vrf_domain_af_control.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("nxos_vrf_domain_af_control.test", "af", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_vrf_domain_af_control.test", "rt_type", "ipv4-ucast"),
				),
			},
			{
				ResourceName:  "nxos_vrf_domain_af_control.test",
				ImportState:   true,
				ImportStateId: "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]",
			},
		},
	})
}

func testAccNxosVRFDomainAfControlConfig_minimum() string {
	return `
	resource "nxos_vrf_domain_af_control" "test" {
		vrf = "VRF1"
		af = "ipv4-ucast"
		rt_type = "ipv4-ucast"
	}
	`
}

func testAccNxosVRFDomainAfControlConfig_all() string {
	return `
	resource "nxos_vrf_domain_af_control" "test" {
		vrf = "VRF1"
		af = "ipv4-ucast"
		rt_type = "ipv4-ucast"
	}
	`
}
