resource "nxos_system" "example" {
  name = "LEAF1"
}

resource "nxos_features" "example" {
  bgp = "enabled"
}

resource "nxos_loopback_interface" "example" {
  interface_id = "lo0"
  admin_state  = "up"
  description  = "Loopback0"
}

resource "nxos_vrf" "example" {
  name        = "VRF1"
  description = "My VRF1 Description"
}
