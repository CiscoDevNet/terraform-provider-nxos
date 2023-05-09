data "nxos_vrf_address_family" "example" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
}
