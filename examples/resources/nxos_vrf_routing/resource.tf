resource "nxos_vrf_routing" "example" {
  vrf                 = "VRF1"
  route_distinguisher = "rd:unknown:0:0"
}
