resource "nxos_evpn_vni_route_target" "example" {
  encap = "vxlan-123456"
  direction = "import"
  route_target = "route-target:as2-nn2:2:2"
}
