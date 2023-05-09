resource "nxos_bgp_peer_template_max_prefix" "example" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  action = "log"
  maximum_prefix = 10000
  restart_time = 0
  threshold = 30
}
