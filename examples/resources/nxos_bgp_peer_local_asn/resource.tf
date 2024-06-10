resource "nxos_bgp_peer_local_asn" "example" {
  asn_propagation = "no-prepend"
  local_asn       = "65001"
  vrf             = "default"
  address         = "192.168.0.1"
}
