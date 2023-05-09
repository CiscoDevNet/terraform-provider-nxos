resource "nxos_evpn_vni" "example" {
  encap               = "vxlan-123456"
  route_distinguisher = "rd:unknown:0:0"
}
