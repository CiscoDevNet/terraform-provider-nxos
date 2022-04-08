data "nxos_bgp_address_family" "example" {
  vrf            = "default"
  address_family = "ipv4-ucast"
}
