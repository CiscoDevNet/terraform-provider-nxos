resource "nxos_bgp_vrf" "example" {
  asn = "65001"
  name = "default"
  router_id = "1.1.1.1"
}
