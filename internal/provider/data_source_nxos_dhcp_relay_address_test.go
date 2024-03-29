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

func TestAccDataSourceNxosDHCPRelayAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDHCPRelayAddressPrerequisitesConfig + testAccDataSourceNxosDHCPRelayAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_address.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_address.test", "address", "1.1.1.1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDHCPRelayAddressPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/dhcp"
  class_name = "fmDhcp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/dhcp/inst/relayif-[eth1/10]"
  class_name = "dhcpRelayIf"
  content = {
      id = "eth1/10"
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosDHCPRelayAddressConfig = `

resource "nxos_dhcp_relay_address" "test" {
  interface_id = "eth1/10"
  vrf = "VRF1"
  address = "1.1.1.1"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_dhcp_relay_address" "test" {
  interface_id = "eth1/10"
  vrf = "VRF1"
  address = "1.1.1.1"
  depends_on = [nxos_dhcp_relay_address.test]
}
`
