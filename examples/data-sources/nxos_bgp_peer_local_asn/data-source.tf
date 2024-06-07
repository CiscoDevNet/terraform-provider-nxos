data "nxos_bgp_peer_local_asn" "example" {
  vrf     = "default"
  address = "192.168.0.1"
}
