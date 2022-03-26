data "nxos_queuing_qos_policy_map_match_class_map_priority" "example" {
  policy_map_name = "PM1"
  class_map_name  = "c-out-q1"
}
