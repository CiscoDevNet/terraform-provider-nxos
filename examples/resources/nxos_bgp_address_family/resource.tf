resource "nxos_bgp_address_family" "example" {
  asn                                    = "65001"
  vrf                                    = "default"
  address_family                         = "ipv4-ucast"
  critical_nexthop_timeout               = "2500"
  non_critical_nexthop_timeout           = "8000"
  advertise_l2vpn_evpn                   = "disabled"
  advertise_physical_ip_for_type5_routes = "disabled"
  max_ecmp_paths                         = 2
  max_external_ecmp_paths                = 1
  max_external_internal_ecmp_paths       = 1
  max_local_ecmp_paths                   = 1
  max_mixed_ecmp_paths                   = 1
  default_information_originate          = "disabled"
  next_hop_route_map_name                = "ROUTEMAP1"
  prefix_priority                        = "none"
  retain_rt_all                          = "disabled"
  advertise_only_active_routes           = "disabled"
  table_map_route_map_name               = "ROUTE_MAP1"
  vni_ethernet_tag                       = "disabled"
  wait_igp_converged                     = "disabled"
}
