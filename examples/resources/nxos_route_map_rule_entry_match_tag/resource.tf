resource "nxos_route_map_rule_entry_match_tag" "example" {
  rule_name = "RULE1"
  order     = 10
  tag       = 12345
}
