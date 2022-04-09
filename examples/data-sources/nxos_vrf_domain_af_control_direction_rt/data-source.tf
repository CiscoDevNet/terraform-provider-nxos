data "nxos_vrf_domain_af_control_direction_rt" "example" {
  vrf       = "VRF1"
  af        = "ipv4-ucast"
  rt_type   = "ipv4-ucast"
  direction = "import"
  rt        = "route-target:as2-nn2:2:2"
}
