resource "nxos_queuing_qos_policy_map" "example" {
  name = "PM1"
  match_type = "match-any"
}
