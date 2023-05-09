data "nxos_route_map_rule_entry_set_regular_community_item" "example" {
  rule_name = "RULE1"
  order     = 10
  community = "regular:as2-nn2:65001:123"
}
