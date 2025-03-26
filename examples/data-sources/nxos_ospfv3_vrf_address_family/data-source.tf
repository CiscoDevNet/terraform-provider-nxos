data "nxos_ospfv3_vrf_address_family" "example" {
  instance_name       = "OSPFv3"
  vrf_name            = "VRF1"
  address_family_type = "ipv6-ucast"
}
