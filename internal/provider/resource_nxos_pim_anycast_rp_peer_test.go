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

func TestAccNxosPIMAnycastRPPeer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMAnycastRPPeerPrerequisitesConfig + testAccNxosPIMAnycastRPPeerConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp_peer.test", "vrf_name", "default"),
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp_peer.test", "address", "10.1.1.1/32"),
					resource.TestCheckResourceAttr("nxos_pim_anycast_rp_peer.test", "rp_set_address", "20.1.1.1/32"),
				),
			},
			{
				ResourceName:  "nxos_pim_anycast_rp_peer.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst/dom-[default]/acastrpfunc/peer-[10.1.1.1/32]-peer-[20.1.1.1/32]",
			},
		},
	})
}

const testAccNxosPIMAnycastRPPeerPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/pim/inst/dom-[default]/acastrpfunc"
  class_name = "pimAcastRPFuncP"
  depends_on = [nxos_rest.PreReq3, ]
}

`

func testAccNxosPIMAnycastRPPeerConfig_minimum() string {
	return `
	resource "nxos_pim_anycast_rp_peer" "test" {
		vrf_name = "default"
		address = "10.1.1.1/32"
		rp_set_address = "20.1.1.1/32"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}

func testAccNxosPIMAnycastRPPeerConfig_all() string {
	return `
	resource "nxos_pim_anycast_rp_peer" "test" {
		vrf_name = "default"
		address = "10.1.1.1/32"
		rp_set_address = "20.1.1.1/32"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}
