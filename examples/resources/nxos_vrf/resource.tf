resource "nxos_vrf" "example" {
  vrfs = {
    "VRF1" = {
      description         = "My VRF1 Description"
      admin_state         = "admin-up"
      controller_id       = 1
      encap               = "vxlan-103901"
      l3vni               = false
      routing_encap       = "unknown"
      route_distinguisher = "rd:unknown:0:0"
      address_families = {
        "ipv4-ucast" = {
          route_target_address_families = {
            "ipv4-ucast" = {
              route_target_directions = {
                "import" = {
                  route_map                = "ABC"
                  advertise_vpn            = true
                  allow_vpn                = false
                  default_vrf_prefix_limit = 1000
                  default_vrf_route_map    = "DEF-RM"
                  route_targets = {
                    "route-target:as2-nn2:2:2" = {}
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
