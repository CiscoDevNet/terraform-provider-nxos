data "nxos_default_qos_policy_map_match_class_map_set_qos_group" "example" {
  policy_map_name = "PM1"
  class_map_name  = "Voice"
}
