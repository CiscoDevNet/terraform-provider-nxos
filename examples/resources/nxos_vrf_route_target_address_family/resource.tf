resource "nxos_vrf_route_target_address_family" "example" {
  vrf                         = "VRF1"
  address_family              = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
}
