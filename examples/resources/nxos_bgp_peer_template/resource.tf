resource "nxos_bgp_peer_template" "example" {
  asn = "65001"
  template_name = "SPINE-PEERS"
  remote_asn = "65002"
  description = "My Description"
  peer_type = "fabric-internal"
  source_interface = "lo0"
}
