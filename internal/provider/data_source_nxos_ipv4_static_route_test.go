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

func TestAccDataSourceNxosIPv4StaticRoute(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosIPv4StaticRoutePrerequisitesConfig + testAccDataSourceNxosIPv4StaticRouteConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "prefix", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.interface_id", "unspecified"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.address", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.vrf_name", "default"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.object", "10"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.preference", "123"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_static_route.test", "next_hops.0.tag", "10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4StaticRoutePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipv4/inst/dom-[default]"
  class_name = "ipv4Dom"
}

`

const testAccDataSourceNxosIPv4StaticRouteConfig = `

resource "nxos_ipv4_static_route" "test" {
  vrf_name = "default"
  prefix = "1.1.1.0/24"
  next_hops = [{
    interface_id = "unspecified"
    address = "1.2.3.4"
    vrf_name = "default"
    description = "My Description"
    object = 10
    preference = 123
    tag = 10
  }]
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_ipv4_static_route" "test" {
  vrf_name = "default"
  prefix = "1.1.1.0/24"
  depends_on = [nxos_ipv4_static_route.test]
}
`
