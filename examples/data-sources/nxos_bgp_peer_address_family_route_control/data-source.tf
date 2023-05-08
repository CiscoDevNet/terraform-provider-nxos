data "nxos_bgp_peer_address_family_route_control" "example" {
  asn            = "65001"
  vrf            = "default"
  address        = "192.168.0.1"
  address_family = "ipv4-ucast"
  direction      = "in"
}
