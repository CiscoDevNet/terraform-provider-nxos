resource "nxos_bgp_instance_domain_af" "example" {
  vrf                          = "default"
  af                           = "ipv4-ucast"
  critical_nexthop_timeout     = 1800
  non_critical_nexthop_timeout = 1800
}
