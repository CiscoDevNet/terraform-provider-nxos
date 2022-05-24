data "nxos_bgp_graceful_restart" "example" {
  asn = "65001"
  vrf = "default"
}
