resource "nxos_evpn" "example" {
  admin_state = "enabled"
  vnis = [{
    encap               = "vxlan-123456"
    route_distinguisher = "rd:unknown:0:0"
    table_map           = "ROUTE_MAP1"
    table_map_filter    = false
    route_target_directions = [{
      type = "import"
      route_targets = [{
        route_target = "route-target:as2-nn2:2:2"
      }]
    }]
  }]
}
