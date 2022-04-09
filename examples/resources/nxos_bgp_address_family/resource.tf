resource "nxos_bgp_address_family" "example" {
  vrf                          = "default"
  address_family               = "ipv4-ucast"
  critical_nexthop_timeout     = 1800
  non_critical_nexthop_timeout = 1800
}
