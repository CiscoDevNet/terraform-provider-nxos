resource "nxos_bgp_peer_template_max_prefix" "example" {
  template_name  = "SPINE-PEERS"
  address_family = "ipv4-ucast"
  action         = "log"
  maximum_prefix = 10000
  restart_time   = 0
  threshold      = 30
}
