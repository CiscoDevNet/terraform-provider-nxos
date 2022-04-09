resource "nxos_bgp_instance_domain_peercontainer_af_maxprefix" "example" {
  vrf            = "default"
  template_name  = "SPINE-PEERS"
  af             = "ipv4-ucast"
  action         = "log"
  maximum_prefix = 10000
  restart_time   = 0
  threshold      = 30
}
