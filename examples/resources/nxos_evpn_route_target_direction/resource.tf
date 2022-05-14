resource "nxos_evpn_route_target_direction" "example" {
  encap     = "vxlan-123456"
  direction = "import"
}
