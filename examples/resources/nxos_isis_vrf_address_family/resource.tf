resource "nxos_isis_vrf_address_family" "example" {
  instance_name               = "ISIS1"
  vrf                         = "default"
  address_family              = "v4"
  segment_routing_mpls        = true
  enable_bfd                  = false
  prefix_advertise_passive_l1 = true
  prefix_advertise_passive_l2 = true
}
