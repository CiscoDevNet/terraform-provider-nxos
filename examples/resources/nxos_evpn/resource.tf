resource "nxos_evpn" "example" {
  admin_state = "enabled"
  vnis = {
    "vxlan-123456" = {
      route_distinguisher = "rd:unknown:0:0"
      table_map = "ROUTE_MAP1"
      table_map_filter = false
      route_target_directions = {
        "import" = {
          route_targets = {
            "route-target:as2-nn2:2:2" = {}
          }
        }
      }
    }
  }
}
