data "nxos_bgp_peer_address_family" "example" {
  vrf            = "default"
  address        = "192.168.0.1"
  address_family = "ipv4-ucast"
}
