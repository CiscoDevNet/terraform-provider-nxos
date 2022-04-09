resource "nxos_bgp_vrf" "example" {
  name      = "default"
  router_id = "1.1.1.1"
}
