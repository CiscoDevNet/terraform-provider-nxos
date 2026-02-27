resource "nxos_vrf" "example" {
  name                = "VRF1"
  description         = "My VRF1 Description"
  admin_state         = "admin-up"
  controller_id       = 1
  encap               = "vxlan-103901"
  l3vni               = false
  oui                 = "000000"
  vpn_id              = "100:200"
  routing_encap       = "unknown"
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
