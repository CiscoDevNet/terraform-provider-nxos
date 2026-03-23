resource "nxos_system" "example" {
  name = "LEAF1"
}

resource "nxos_feature" "example" {
  bgp = "enabled"
}

resource "nxos_loopback_interface" "example" {
  loopback_interfaces = {
    "lo0" = {
      admin_state = "up"
      description = "Loopback0"
    }
  }
}

resource "nxos_vrf" "example" {
  vrfs = {
    "VRF1" = {
      description = "My VRF1 Description"
    }
  }
}
