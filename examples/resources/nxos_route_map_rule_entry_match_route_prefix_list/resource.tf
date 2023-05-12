resource "nxos_route_map_rule_entry_match_route_prefix_list" "example" {
  rule_name      = "RULE1"
  order          = 10
  prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
}
