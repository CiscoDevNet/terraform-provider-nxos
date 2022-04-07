data "nxos_bgp_instance_domain_peer" "example" {
  vrf     = "default"
  address = "192.168.0.1"
}
