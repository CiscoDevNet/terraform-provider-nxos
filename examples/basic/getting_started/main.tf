resource "nxos_system" "example" {
  name = "LEAF1"
}

resource "nxos_feature_bgp" "example" {
  admin_state = "enabled"
}

resource "nxos_loopback_interface" "example" {
  interface_id = "lo0"
  description  = "Loopback0"
}

resource "nxos_vrf" "example" {
  name        = "VRF1"
  description = "My VRF1 Description"
}
