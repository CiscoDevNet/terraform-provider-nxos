data "nxos_bgp_instance_domain_peer_container_af" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
  af            = "ipv4-ucast"
}
