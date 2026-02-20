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

package provider

import (
	"fmt"
	"testing"

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccNxosRest(t *testing.T) {
	var tfVersion *goversion.Version
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			terraformVersionCapture{Version: &tfVersion},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNxosRestConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "id", "sys/intf/phys-[eth1/1]"),
				),
			},
			{
				Config: testAccNxosRestConfig_child(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "class_name", "ipqosCMapInst"),
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "id", "sys/ipqos/dflt/c/name-[CM1]"),
					resource.TestCheckResourceAttr("nxos_rest.ipqosCMapInst", "children.0.content.val", "ef"),
				),
			},
			{
				Config: testAccNxosRestConfig_interface("Create description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "class_name", "l1PhysIf"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "id", "sys/intf/phys-[eth1/1]"),
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "content.descr", "Create description"),
				),
			},
			{
				ResourceName:  "nxos_rest.l1PhysIf",
				ImportState:   true,
				ImportStateId: "sys/intf/phys-[eth1/1],l1PhysIf",
			},
			{
				ResourceName:       "nxos_rest.l1PhysIf",
				ImportState:        true,
				ImportStateKind:    resource.ImportBlockWithResourceIdentity,
				ExpectNonEmptyPlan: true,
				SkipFunc:           skipBelowTerraformVersion(&tfVersion, goversion.Must(goversion.NewVersion("1.12.0"))),
			},
			{
				Config: testAccNxosRestConfig_interface("Updated description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_rest.l1PhysIf", "content.descr", "Updated description"),
				),
			},
		},
	})
}

func testAccNxosRestConfig_empty() string {
	return `
	resource "nxos_rest" "l1PhysIf" {
		dn = "sys/intf/phys-[eth1/1]"
		class_name = "l1PhysIf"
	}
	`
}

func testAccNxosRestConfig_interface(description string) string {
	return fmt.Sprintf(`
	resource "nxos_rest" "l1PhysIf" {
		dn = "sys/intf/phys-[eth1/1]"
		class_name = "l1PhysIf"
		content = {
			id = "eth1/1"
			descr = "%[1]s"
		}
	}
	`, description)
}

func testAccNxosRestConfig_child() string {
	return `
	resource "nxos_rest" "ipqosCMapInst" {
		dn = "sys/ipqos/dflt/c/name-[CM1]"
		class_name = "ipqosCMapInst"
		content = {
			name = "CM1"
		}
		children = [
            {
                rn = "dscp-ef"
                class_name = "ipqosDscp"
				content = {
					val = "ef"
				}
			}
		]
	}
	`
}
