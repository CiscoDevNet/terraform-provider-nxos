data "nxos_default_qos_policy_map_match_class_map" "example" {
  policy_map_name = "PM1"
  name            = "Voice"
}
