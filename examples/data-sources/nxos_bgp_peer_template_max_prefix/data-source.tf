data "nxos_bgp_peer_template_max_prefix" "example" {
  template_name  = "SPINE-PEERS"
  address_family = "ipv4-ucast"
}
