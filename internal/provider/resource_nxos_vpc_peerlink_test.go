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

func TestAccNxosVPCPeerlink(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVPCPeerlinkPrerequisitesConfig + testAccNxosVPCPeerlinkConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vpc_peerlink.test", "port_channel_id", "po1"),
				),
			},
			{
				ResourceName:  "nxos_vpc_peerlink.test",
				ImportState:   true,
				ImportStateId: "sys/vpc/inst/dom/keepalive/peerlink",
			},
		},
	})
}

const testAccNxosVPCPeerlinkPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/vpc"
  class_name = "fmVpc"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/vpc/inst"
  class_name = "vpcInst"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/vpc/inst/dom"
  class_name = "vpcDom"
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/vpc/inst/dom/keepalive"
  class_name = "vpcKeepalive"
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/fm/lacp"
  class_name = "fmLacp"
  content = {
       = ""
  }
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/intf/aggr-[123]"
  class_name = "pcAggrIf"
  content = {
       = ""
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

func testAccNxosVPCPeerlinkConfig_minimum() string {
	return `
	resource "nxos_vpc_peerlink" "test" {
		port_channel_id = "po1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

func testAccNxosVPCPeerlinkConfig_all() string {
	return `
	resource "nxos_vpc_peerlink" "test" {
		port_channel_id = "po1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}
