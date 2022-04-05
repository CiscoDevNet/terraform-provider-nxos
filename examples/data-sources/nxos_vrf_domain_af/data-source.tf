data "nxos_vrf_domain_af" "example" {
  vrf = "VRF1"
  af  = "ipv4-ucast"
}
