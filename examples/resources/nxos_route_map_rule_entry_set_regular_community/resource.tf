resource "nxos_route_map_rule_entry_set_regular_community" "example" {
  rule_name    = "RULE1"
  order        = 10
  additive     = "disabled"
  no_community = "disabled"
  set_criteria = "none"
}
