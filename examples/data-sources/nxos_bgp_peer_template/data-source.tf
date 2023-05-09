data "nxos_bgp_peer_template" "example" {
  asn = "65001"
  template_name = "SPINE-PEERS"
}
