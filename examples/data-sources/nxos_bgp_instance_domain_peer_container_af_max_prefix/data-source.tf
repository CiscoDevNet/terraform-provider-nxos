data "nxos_bgp_instance_domain_peer_container_af_max_prefix" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
  af            = "ipv4-ucast"
}
