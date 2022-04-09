data "nxos_bgp_instance_domain_af" "example" {
  vrf = "default"
  af  = "ipv4-ucast"
}
