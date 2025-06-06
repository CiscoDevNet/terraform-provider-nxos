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

func TestAccNxosIPv6Interface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv6InterfacePrerequisitesConfig + testAccNxosIPv6InterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "auto_configuration", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "default_route", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "ipv6_forward", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "forward", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "link_address_use_bia", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "use_link_local_address", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "urpf", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv6_interface.test", "link_local_address", "2001:db8:3333:4444:5555:6666:7777:8888"),
				),
			},
			{
				ResourceName:  "nxos_ipv6_interface.test",
				ImportState:   true,
				ImportStateId: "sys/ipv6/inst/dom-[default]/if-[eth1/10]",
			},
		},
	})
}

const testAccNxosIPv6InterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipv6/inst/dom-[default]"
  class_name = "ipv6Dom"
}

`

func testAccNxosIPv6InterfaceConfig_minimum() string {
	return `
	resource "nxos_ipv6_interface" "test" {
		vrf = "default"
		interface_id = "eth1/10"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosIPv6InterfaceConfig_all() string {
	return `
	resource "nxos_ipv6_interface" "test" {
		vrf = "default"
		interface_id = "eth1/10"
		auto_configuration = "disabled"
		default_route = "disabled"
		ipv6_forward = "disabled"
		forward = "disabled"
		link_address_use_bia = "disabled"
		use_link_local_address = "disabled"
		urpf = "disabled"
		link_local_address = "2001:db8:3333:4444:5555:6666:7777:8888"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}
