resource "nxos_bgp_address_family" "example" {
  asn                          = "65001"
  vrf                          = "default"
  address_family               = "ipv4-ucast"
  critical_nexthop_timeout     = 1800
  non_critical_nexthop_timeout = 1800
  adv_l2vpn_evpn               = "disabled"
  adv_phyip_for_type5_routes   = "disabled"
  max_ecmp_paths               = 2
  max_ext_ecmp_paths           = 1
  max_ext_int_ecmp_paths       = 1
  max_eqcost_ecmp_paths        = 1
  max_mixcost_ecmp_paths       = 1
  def_inf_originate            = "disabled"
  next_hop_route_map_name      = "ROUTEMAP1"
  prefix_priority              = "none"
  retain_rt_all                = "disabled"
  adv_only_act_routes          = "disabled"
  route_map_name               = "ROUTE_MAP1"
  vni_ethtag                   = "disabled"
  wait_igp_conv                = "disabled"
}
