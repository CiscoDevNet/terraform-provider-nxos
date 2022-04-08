data "nxos_bgp_peer_template_address_family" "example" {
  vrf            = "default"
  template_name  = "SPINE-PEERS"
  address_family = "ipv4-ucast"
}
