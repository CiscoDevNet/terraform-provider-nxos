resource "nxos_bgp_instance_domain" "example" {
  vrf       = "default"
  router_id = "1.1.1.1"
}
