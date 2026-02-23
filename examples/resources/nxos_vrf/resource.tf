resource "nxos_vrf" "example" {
  name                = "VRF1"
  description         = "My VRF1 Description"
  encap               = "vxlan-103901"
  route_distinguisher = "rd:unknown:0:0"
  address_families = [{
    address_family = "ipv4-ucast"
    route_target_address_families = [{
      route_target_address_family = "ipv4-ucast"
      route_target_directions = [{
        direction = "import"
        route_targets = [{
          route_target = "route-target:as2-nn2:2:2"
        }]
      }]
    }]
  }]
}
