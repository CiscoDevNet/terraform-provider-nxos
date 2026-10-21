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
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// A BGP peer added to an existing BGP configuration, reaching it through the existing instance and default VRF.
var testDeletionPolicyBGP = BGP{
	Asn: types.StringValue("65001"),
	Vrfs: map[string]BGPVrfs{
		"default": {
			Peers: map[string]BGPVrfsPeers{
				"10.0.0.5": {RemoteAsn: types.StringValue("65002")},
			},
		},
	},
}

func TestNxos_DeletionPolicyResetBody(t *testing.T) {
	cases := []struct {
		name        string
		preexisting map[string]bool
		expected    string
	}{
		{
			// The created peer is deleted, the existing instance and VRF stay and the ASN identifying the instance is kept
			name:        "preexisting parents",
			preexisting: map[string]bool{"sys/bgp": true, "sys/bgp/inst": true, "sys/bgp/inst/dom-default": true},
			expected:    `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"children":[{"bgpDom":{"attributes":{"rn":"dom-default"},"children":[{"bgpPeer":{"attributes":{"rn":"peer-[10.0.0.5]","status":"deleted"}}}]}}],"attributes":{"rn":"inst"}}}]}}`,
		},
		{
			// Everything was created by the resource, so it is deleted like with the delete policy
			name:        "created",
			preexisting: map[string]bool{},
			expected:    `{"bgpEntity":{"attributes":{"status":"deleted"}}}`,
		},
		{
			// Without a record all objects are treated as preexisting and are only reset
			name:        "not recorded",
			preexisting: nil,
			expected:    `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"children":[{"bgpDom":{"attributes":{"rn":"dom-default"},"children":[{"bgpPeer":{"attributes":{"rn":"peer-[10.0.0.5]","asn":"DME_UNSET_PROPERTY_MARKER"}}}]}}],"attributes":{"rn":"inst"}}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := testDeletionPolicyBGP.toResetBody(c.preexisting).Str; out != c.expected {
				t.Fatalf("Error matching output and expected:\n%s\nvs\n%s", out, c.expected)
			}
		})
	}
}

// Single child objects are only deleted when the resource configures them. Deleting unconfigured ones removes
// config the resource does not manage and fails for feature-gated classes (e.g. ngoamEntity without feature ngoam).
func TestNxos_DeletionPolicyDeleteBody(t *testing.T) {
	cases := []struct {
		name     string
		data     System
		expected string
	}{
		{
			name:     "unconfigured children",
			data:     System{Name: types.StringValue("LEAF1")},
			expected: `{"topSystem":{"attributes":{"name":"DME_UNSET_PROPERTY_MARKER"}}}`,
		},
		{
			name:     "configured child",
			data:     System{NgoamInstallAcl: types.BoolValue(true)},
			expected: `{"topSystem":{"attributes":{},"children":[{"ngoamEntity":{"attributes":{},"children":[{"ngoamInstallAcl":{"attributes":{"rn":"installacl","status":"deleted"}}}]}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if out := c.data.toDeleteBody().Str; out != c.expected {
				t.Fatalf("Error matching output and expected:\n%s\nvs\n%s", out, c.expected)
			}
		})
	}
}

func TestNxos_DeletionPolicyUpdateBody(t *testing.T) {
	vrf := func(v BGPVrfs) BGP {
		return BGP{Asn: types.StringValue("65001"), Vrfs: map[string]BGPVrfs{"default": v}}
	}
	state := vrf(BGPVrfs{GracefulRestartInterval: types.Int64Value(240), GracefulRestartStaleInterval: types.Int64Value(600)})
	deleteAll := func(string) bool { return true }
	keepAll := func(string) bool { return false }

	cases := []struct {
		name          string
		plan          BGP
		deleteEmptied func(dn string) bool
		checked       []string
		expected      string
	}{
		{
			name:          "emptied with delete policy",
			plan:          vrf(BGPVrfs{}),
			deleteEmptied: deleteAll,
			checked:       []string{"sys/bgp/inst/dom-default/gr"},
			expected:      `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"attributes":{"asn":"65001"},"children":[{"bgpDom":{"attributes":{"name":"default"},"children":[{"bgpGr":{"attributes":{"status":"deleted"}}}]}}]}}]}}`,
		},
		{
			name:          "emptied with unmanaged entries below",
			plan:          vrf(BGPVrfs{}),
			deleteEmptied: keepAll,
			checked:       []string{"sys/bgp/inst/dom-default/gr"},
			expected:      `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"attributes":{"asn":"65001"},"children":[{"bgpDom":{"attributes":{"name":"default"},"children":[{"bgpGr":{"attributes":{"restartIntvl":"DME_UNSET_PROPERTY_MARKER","staleIntvl":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}]}}`,
		},
		{
			name:          "emptied with reset or retain policy",
			plan:          vrf(BGPVrfs{}),
			deleteEmptied: nil,
			expected:      `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"attributes":{"asn":"65001"},"children":[{"bgpDom":{"attributes":{"name":"default"},"children":[{"bgpGr":{"attributes":{"restartIntvl":"DME_UNSET_PROPERTY_MARKER","staleIntvl":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}]}}`,
		},
		{
			name:          "partially removed",
			plan:          vrf(BGPVrfs{GracefulRestartInterval: types.Int64Value(240)}),
			deleteEmptied: deleteAll,
			expected:      `{"bgpEntity":{"attributes":{},"children":[{"bgpInst":{"attributes":{"asn":"65001"},"children":[{"bgpDom":{"attributes":{"name":"default"},"children":[{"bgpGr":{"attributes":{"restartIntvl":"240","staleIntvl":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var checked []string
			var deleteEmptied func(dn string) bool
			if c.deleteEmptied != nil {
				deleteEmptied = func(dn string) bool { checked = append(checked, dn); return c.deleteEmptied(dn) }
			}
			if out := c.plan.toBodyWithDeletes(context.Background(), state, c.plan, false, deleteEmptied).Str; out != c.expected {
				t.Fatalf("Error matching output and expected:\n%s\nvs\n%s", out, c.expected)
			}
			if fmt.Sprint(checked) != fmt.Sprint(c.checked) {
				t.Fatalf("Error matching checked objects and expected: %v vs %v", checked, c.checked)
			}
		})
	}
}

// An ethernet segment, an atomic object that cannot be reset, configured on an existing port-channel.
var testDeletionPolicyAtomic = PortChannelInterface{
	PortChannelInterfaces: map[string]PortChannelInterfacePortChannelInterfaces{
		"po1": {EthernetSegmentType: types.StringValue("native"), EthernetSegmentIdentifier: types.StringValue("0011.2233.4455.6677.8899")},
	},
}

func TestNxos_DeletionPolicyAtomicResetBody(t *testing.T) {
	cases := []struct {
		name        string
		preexisting map[string]bool
		kept        []string
		expected    string
	}{
		{
			// The ethernet segment created by the resource is deleted, the existing port-channel stays
			name:        "created",
			preexisting: map[string]bool{"sys/intf": true, "sys/intf/aggr-[po1]": true},
			expected:    `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"rn":"aggr-[po1]"},"children":[{"nvoEtherSeg":{"attributes":{"rn":"ethernetsegment","status":"deleted"}}}]}}]}}`,
		},
		{
			// The ethernet segment existed before, it cannot be reset and is left unchanged
			name:        "preexisting",
			preexisting: map[string]bool{"sys/intf": true, "sys/intf/aggr-[po1]": true, "sys/intf/aggr-[po1]/ethernetsegment": true},
			kept:        []string{"sys/intf/aggr-[po1]/ethernetsegment"},
			expected:    `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"rn":"aggr-[po1]"}}}]}}`,
		},
		{
			// Without a record all objects are treated as preexisting
			name:        "not recorded",
			preexisting: nil,
			kept:        []string{"sys/intf/aggr-[po1]/ethernetsegment"},
			expected:    `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"rn":"aggr-[po1]"}}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var kept []string
			if out := testDeletionPolicyAtomic.toResetBody(c.preexisting, func(dn string) { kept = append(kept, dn) }).Str; out != c.expected {
				t.Fatalf("Error matching output and expected:\n%s\nvs\n%s", out, c.expected)
			}
			if fmt.Sprint(kept) != fmt.Sprint(c.kept) {
				t.Fatalf("Error matching kept objects and expected: %v vs %v", kept, c.kept)
			}
		})
	}
}

func TestNxos_DeletionPolicyAtomicUpdateBody(t *testing.T) {
	emptied := PortChannelInterface{PortChannelInterfaces: map[string]PortChannelInterfacePortChannelInterfaces{"po1": {}}}
	changed := PortChannelInterface{PortChannelInterfaces: map[string]PortChannelInterfacePortChannelInterfaces{
		"po1": {EthernetSegmentType: types.StringValue("native"), EthernetSegmentSystemMac: types.StringValue("00:00:11:11:22:22")},
	}}
	deleteAll := func(string) bool { return true }
	keepAll := func(string) bool { return false }

	cases := []struct {
		name          string
		plan          PortChannelInterface
		deleteEmptied func(dn string) bool
		atomicEmptied func(dn string) bool
		checked       []string
		expected      string
	}{
		{
			name:          "emptied with delete policy",
			plan:          emptied,
			deleteEmptied: deleteAll,
			checked:       []string{"sys/intf/aggr-[po1]/ethernetsegment"},
			expected:      `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"id":"po1"},"children":[{"nvoEtherSeg":{"attributes":{"status":"deleted"}}}]}}]}}`,
		},
		{
			name:          "emptied and created by the resource with reset or retain policy",
			plan:          emptied,
			atomicEmptied: deleteAll,
			checked:       []string{"sys/intf/aggr-[po1]/ethernetsegment"},
			expected:      `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"id":"po1"},"children":[{"nvoEtherSeg":{"attributes":{"status":"deleted"}}}]}}]}}`,
		},
		{
			name:          "emptied and preexisting with reset or retain policy",
			plan:          emptied,
			atomicEmptied: keepAll,
			checked:       []string{"sys/intf/aggr-[po1]/ethernetsegment"},
			expected:      `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"id":"po1"}}}]}}`,
		},
		{
			name:          "partially removed",
			plan:          changed,
			atomicEmptied: deleteAll,
			expected:      `{"interfaceEntity":{"attributes":{},"children":[{"pcAggrIf":{"attributes":{"id":"po1"},"children":[{"nvoEtherSeg":{"attributes":{"esType":"native","sysMac":"00:00:11:11:22:22","esi":"DME_UNSET_PROPERTY_MARKER"}}}]}}]}}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var checked []string
			record := func(f func(dn string) bool) func(dn string) bool {
				if f == nil {
					return nil
				}
				return func(dn string) bool { checked = append(checked, dn); return f(dn) }
			}
			if out := c.plan.toBodyWithDeletes(context.Background(), testDeletionPolicyAtomic, c.plan, false, record(c.deleteEmptied), record(c.atomicEmptied)).Str; out != c.expected {
				t.Fatalf("Error matching output and expected:\n%s\nvs\n%s", out, c.expected)
			}
			if fmt.Sprint(checked) != fmt.Sprint(c.checked) {
				t.Fatalf("Error matching checked objects and expected: %v vs %v", checked, c.checked)
			}
		})
	}
}

const testAccDeletionPolicyUnmanagedConfig = `
resource "nxos_dme" "unmanaged" {
  dn         = "sys/rpm/pfxlistv4-[PL_UNMANAGED]"
  class_name = "rtpfxRuleV4"
  delete     = %s
  content = {
    name  = "PL_UNMANAGED"
    descr = "unmanaged"
  }
}
`

func testAccDeletionPolicyUnmanaged(delete string) string {
	return fmt.Sprintf(testAccDeletionPolicyUnmanagedConfig, delete)
}

func testAccDeletionPolicyManagedConfig(policy string) string {
	config := `
resource "nxos_route_policy" "test" {
`
	if policy != "" {
		config += `  deletion_policy = "` + policy + `"
`
	}
	return config + `  ipv4_prefix_lists = {
    "PL_MANAGED" = {
      description = "managed"
    }
  }
}
`
}

const testAccDeletionPolicyDataSourceConfig = `
data "nxos_route_policy" "test" {}
`

// TestAccRoutePolicyResource_DeletionPolicyReset destroys a resource with the reset policy and checks that the list entry it created
// is deleted, while a list entry configured outside of the resource stays.
func TestAccRoutePolicyResource_DeletionPolicyReset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletionPolicyUnmanaged("false") + testAccDeletionPolicyManagedConfig("reset"),
			},
			{
				Config: testAccDeletionPolicyUnmanaged("false"),
			},
			{
				Config: testAccDeletionPolicyUnmanaged("false") + testAccDeletionPolicyDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_route_policy.test", "ipv4_prefix_lists.PL_UNMANAGED.description", "unmanaged"),
					resource.TestCheckNoResourceAttr("data.nxos_route_policy.test", "ipv4_prefix_lists.PL_MANAGED.description"),
				),
			},
			{
				Config: testAccDeletionPolicyUnmanaged("true"),
			},
		},
	})
}

// TestAccRoutePolicyResource_DeletionPolicyRetain destroys a resource with the retain policy and checks that its configuration stays
// on the device.
func TestAccRoutePolicyResource_DeletionPolicyRetain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletionPolicyManagedConfig("retain"),
			},
			{
				Config: testAccDeletionPolicyDataSourceConfig,
			},
			{
				Config: testAccDeletionPolicyDataSourceConfig,
				Check:  resource.TestCheckResourceAttr("data.nxos_route_policy.test", "ipv4_prefix_lists.PL_MANAGED.description", "managed"),
			},
			{
				// Manage the retained configuration again to remove it at the end of the test
				Config: testAccDeletionPolicyManagedConfig(""),
			},
		},
	})
}

// TestAccRoutePolicyResource_ProviderDeletionPolicy sets the retain policy for all resources with the environment variable of the
// provider, and checks that the configuration of a resource without a policy of its own stays on the device.
func TestAccRoutePolicyResource_ProviderDeletionPolicy(t *testing.T) {
	t.Setenv("NXOS_DELETION_POLICY", "retain")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletionPolicyManagedConfig(""),
			},
			{
				Config: testAccDeletionPolicyDataSourceConfig,
			},
			{
				Config: testAccDeletionPolicyDataSourceConfig,
				Check:  resource.TestCheckResourceAttr("data.nxos_route_policy.test", "ipv4_prefix_lists.PL_MANAGED.description", "managed"),
			},
			{
				// The policy of the resource overrides the one of the provider and removes the configuration at the end of the test
				Config: testAccDeletionPolicyManagedConfig("delete"),
			},
		},
	})
}

// testAccDeletionPolicyBGPUnmanagedConfig is an existing BGP configuration with a peer that is not managed by Terraform.
const testAccDeletionPolicyBGPUnmanagedConfig = `
resource "nxos_dme" "feature_bgp" {
  dn         = "sys/fm/bgp"
  class_name = "fmBgp"
  delete     = false
  content = {
    adminSt = "enabled"
  }
}

resource "nxos_dme" "bgp" {
  dn         = "sys/bgp"
  class_name = "bgpEntity"
  delete     = false
  depends_on = [nxos_dme.feature_bgp]
}

resource "nxos_dme" "instance" {
  dn         = "sys/bgp/inst"
  class_name = "bgpInst"
  delete     = false
  content = {
    asn = "65001"
  }
  depends_on = [nxos_dme.bgp]
}

resource "nxos_dme" "vrf" {
  dn         = "sys/bgp/inst/dom-default"
  class_name = "bgpDom"
  delete     = false
  content = {
    name = "default"
  }
  depends_on = [nxos_dme.instance]
}

resource "nxos_dme" "unmanaged_peer" {
  dn         = "sys/bgp/inst/dom-default/peer-[10.0.0.1]"
  class_name = "bgpPeer"
  delete     = false
  content = {
    addr = "10.0.0.1"
    asn  = "65002"
  }
  depends_on = [nxos_dme.vrf]
}
`

func testAccDeletionPolicyBGPManagedConfig(policy string, peer bool) string {
	config := `
resource "nxos_bgp" "test" {
  deletion_policy = "` + policy + `"
  asn             = "65001"
`
	if peer {
		config += `  vrfs = {
    "default" = {
      peers = {
        "10.0.0.5" = {
          remote_asn = "65004"
        }
      }
    }
  }
`
	}
	return config + `  depends_on = [nxos_dme.unmanaged_peer]
}
`
}

// TestAccBGPResource_DeletionPolicyReset adds a peer to an existing BGP configuration and destroys it with the reset policy,
// which deletes only the added peer, while the instance, the default VRF and the existing peer stay.
func TestAccBGPResource_DeletionPolicyReset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletionPolicyBGPUnmanagedConfig + testAccDeletionPolicyBGPManagedConfig("reset", true),
			},
			{
				Config: testAccDeletionPolicyBGPUnmanagedConfig,
			},
			{
				Config: testAccDeletionPolicyBGPUnmanagedConfig + `data "nxos_bgp" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp.test", "asn", "65001"),
					resource.TestCheckResourceAttr("data.nxos_bgp.test", "vrfs.default.peers.10.0.0.1.remote_asn", "65002"),
					resource.TestCheckNoResourceAttr("data.nxos_bgp.test", "vrfs.default.peers.10.0.0.5.remote_asn"),
				),
			},
			{
				// Manage the BGP configuration with the delete policy to remove it at the end of the test
				Config: testAccDeletionPolicyBGPUnmanagedConfig + testAccDeletionPolicyBGPManagedConfig("delete", false),
			},
		},
	})
}

const testAccDeletionPolicyUnmanagedDescriptionConfig = `
resource "nxos_dme" "unmanaged_description" {
  dn         = "sys/intf/phys-[eth1/20]"
  class_name = "l1PhysIf"
  delete     = false
  content = {
    descr = "unmanaged"
  }
}
`

func testAccDeletionPolicyPhysicalInterfaceConfig(policy string, managed string) string {
	return testAccDeletionPolicyUnmanagedDescriptionConfig + `
resource "nxos_physical_interface" "test" {
  deletion_policy = "` + policy + `"
  physical_interfaces = {
    "eth1/20" = {
      ` + managed + `
    }
  }
  depends_on = [nxos_dme.unmanaged_description]
}
`
}

// TestAccPhysicalInterfaceResource_DeletionPolicyReset destroys a resource whose objects cannot be deleted with the reset
// policy, which resets only the managed attribute of the interface and keeps its unmanaged description.
func TestAccPhysicalInterfaceResource_DeletionPolicyReset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDeletionPolicyPhysicalInterfaceConfig("reset", "delay = 10"),
				Check:  resource.TestCheckResourceAttr("nxos_physical_interface.test", "physical_interfaces.eth1/20.delay", "10"),
			},
			{
				Config: testAccDeletionPolicyUnmanagedDescriptionConfig,
			},
			{
				Config: testAccDeletionPolicyUnmanagedDescriptionConfig + `data "nxos_physical_interface" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_physical_interface.test", "physical_interfaces.eth1/20.description", "unmanaged"),
					resource.TestCheckResourceAttrWith("data.nxos_physical_interface.test", "physical_interfaces.eth1/20.delay", func(value string) error {
						if value == "10" {
							return fmt.Errorf("expected the delay to be reset")
						}
						return nil
					}),
				),
			},
			{
				// Manage the description with the delete policy to reset it at the end of the test
				Config: testAccDeletionPolicyPhysicalInterfaceConfig("delete", `description = "unmanaged"`),
			},
		},
	})
}
