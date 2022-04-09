resource "nxos_bgp_instance_domain_peer_container" "example" {
  vrf              = "default"
  template_name    = "SPINE-PEERS"
  asn              = "65002"
  peer_type        = "fabric-internal"
  source_interface = "lo0"
}
