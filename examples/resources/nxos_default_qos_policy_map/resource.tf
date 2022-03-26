resource "nxos_default_qos_policy_map" "example" {
  name       = "PM1"
  match_type = "match-any"
}
