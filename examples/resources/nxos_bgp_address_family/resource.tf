resource "nxos_bgp_address_family" "example" {
  asn = "65001"
  vrf = "default"
  address_family = "ipv4-ucast"
  critical_nexthop_timeout = 1800
  non_critical_nexthop_timeout = 1800
}
