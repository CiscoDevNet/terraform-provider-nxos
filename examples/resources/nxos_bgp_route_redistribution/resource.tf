resource "nxos_bgp_route_redistribution" "example" {
  asn               = "65001"
  vrf               = "default"
  address_family    = "ipv4-ucast"
  protocol          = "ospf"
  protocol_instance = "OSPF1"
  route_map         = "route_map_ospf_1"
  scope             = "inter"
  srv6_prefix_type  = ""
}
