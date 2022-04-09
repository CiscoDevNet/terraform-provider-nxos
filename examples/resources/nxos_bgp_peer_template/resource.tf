resource "nxos_bgp_peer_template" "example" {
  vrf              = "default"
  template_name    = "SPINE-PEERS"
  asn              = "65002"
  peer_type        = "fabric-internal"
  source_interface = "lo0"
}
