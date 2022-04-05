resource "nxos_vrf_domain" "example" {
  vrf = "VRF1"
  rd  = "rd:unknown:0:0"
}
