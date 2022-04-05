data "nxos_vrf_domain_af_control" "example" {
  vrf     = "VRF1"
  af      = "ipv4-ucast"
  rt_type = "ipv4-ucast"
}
