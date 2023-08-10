resource "nxos_bgp_advertised_prefix" "example" {
  asn            = "65001"
  vrf            = "default"
  address_family = "ipv4-ucast"
  prefix         = "192.168.1.0/24"
  route_map      = "rt-map"
}
