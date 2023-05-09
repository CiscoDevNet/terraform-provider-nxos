data "nxos_bgp_peer_template_max_prefix" "example" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
}
