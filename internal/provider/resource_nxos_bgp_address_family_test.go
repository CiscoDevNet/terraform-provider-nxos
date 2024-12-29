// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNxosBGPAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPAddressFamilyPrerequisitesConfig + testAccNxosBGPAddressFamilyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "asn", "65001"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "critical_nexthop_timeout", "2500"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "non_critical_nexthop_timeout", "8000"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "adv_l2vpn_evpn", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "adv_phyip_for_type5_routes", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "max_ecmp_paths", "2"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "max_ext_ecmp_paths", "1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "max_ext_int_ecmp_paths", "1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "max_eqcost_ecmp_paths", "1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "max_mixcost_ecmp_paths", "1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "def_inf_originate", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "next_hop_route_map_name", "ROUTEMAP1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "prefix_priority", "none"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "retain_rt_all", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "adv_only_act_routes", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "route_map_name", "ROUTE_MAP1"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "vni_ethtag", "disabled"),
					resource.TestCheckResourceAttr("nxos_bgp_address_family.test", "wait_igp_conv", "disabled"),
				),
			},
			{
				ResourceName:  "nxos_bgp_address_family.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/af-[ipv4-ucast]",
			},
		},
	})
}

const testAccNxosBGPAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosBGPAddressFamilyConfig_minimum() string {
	return `
	resource "nxos_bgp_address_family" "test" {
		asn = "65001"
		vrf = "default"
		address_family = "ipv4-ucast"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosBGPAddressFamilyConfig_all() string {
	return `
	resource "nxos_bgp_address_family" "test" {
		asn = "65001"
		vrf = "default"
		address_family = "ipv4-ucast"
		critical_nexthop_timeout = "2500"
		non_critical_nexthop_timeout = "8000"
		adv_l2vpn_evpn = "disabled"
		adv_phyip_for_type5_routes = "disabled"
		max_ecmp_paths = 2
		max_ext_ecmp_paths = 1
		max_ext_int_ecmp_paths = 1
		max_eqcost_ecmp_paths = 1
		max_mixcost_ecmp_paths = 1
		def_inf_originate = "disabled"
		next_hop_route_map_name = "ROUTEMAP1"
		prefix_priority = "none"
		retain_rt_all = "disabled"
		adv_only_act_routes = "disabled"
		route_map_name = "ROUTE_MAP1"
		vni_ethtag = "disabled"
		wait_igp_conv = "disabled"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
