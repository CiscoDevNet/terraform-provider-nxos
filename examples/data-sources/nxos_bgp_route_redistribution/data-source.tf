data "nxos_bgp_route_redistribution" "example" {
  asn               = "65001"
  vrf               = "default"
  address_family    = "ipv4-ucast"
  protocol          = "ospf"
  protocol_instance = "OSPF1"
}
