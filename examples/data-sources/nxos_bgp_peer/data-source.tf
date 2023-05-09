data "nxos_bgp_peer" "example" {
  asn = "65001"
  vrf = "default"
  address = "192.168.0.1"
}
