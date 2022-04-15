resource "nxos_bgp_peer_template" "example" {
  template_name    = "SPINE-PEERS"
  asn              = "65002"
  description      = "My Description"
  peer_type        = "fabric-internal"
  source_interface = "lo0"
}
