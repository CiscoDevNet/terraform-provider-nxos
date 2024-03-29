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

func TestAccDataSourceNxosISISVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosISISVRFPrerequisitesConfig + testAccDataSourceNxosISISVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "authentication_check_l1", "false"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "authentication_check_l2", "false"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "authentication_type_l1", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "authentication_type_l2", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "bandwidth_reference", "400000"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "banwidth_reference_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "is_type", "l2"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "metric_type", "wide"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "mtu", "2000"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "net", "49.0001.0000.0000.3333.00"),
					resource.TestCheckResourceAttr("data.nxos_isis_vrf.test", "passive_default", "l12"),
				),
			},
		},
	})
}

const testAccDataSourceNxosISISVRFPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/isis"
  class_name = "fmIsis"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/isis"
  class_name = "isisEntity"
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/isis/inst-[ISIS1]"
  class_name = "isisInst"
  content = {
      name = "ISIS1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosISISVRFConfig = `

resource "nxos_isis_vrf" "test" {
  instance_name = "ISIS1"
  name = "default"
  admin_state = "enabled"
  authentication_check_l1 = false
  authentication_check_l2 = false
  authentication_key_l1 = ""
  authentication_key_l2 = ""
  authentication_type_l1 = "unknown"
  authentication_type_l2 = "unknown"
  bandwidth_reference = 400000
  banwidth_reference_unit = "mbps"
  is_type = "l2"
  metric_type = "wide"
  mtu = 2000
  net = "49.0001.0000.0000.3333.00"
  passive_default = "l12"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_isis_vrf" "test" {
  instance_name = "ISIS1"
  name = "default"
  depends_on = [nxos_isis_vrf.test]
}
`
