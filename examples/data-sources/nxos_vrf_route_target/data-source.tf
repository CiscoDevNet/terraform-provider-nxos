data "nxos_vrf_route_target" "example" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  route_target_address_family = "ipv4-ucast"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
}
