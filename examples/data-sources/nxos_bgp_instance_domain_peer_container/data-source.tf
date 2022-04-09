data "nxos_bgp_instance_domain_peer_container" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
}
