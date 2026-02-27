resource "nxos_queuing_qos" "example" {
  policy_maps = [{
    name       = "PM1"
    match_type = "match-any"
    match_class_maps = [{
      name                = "c-out-q1"
      next_class_map      = "c-out-q2"
      previous_class_map  = "c-out-q2"
      priority            = 1
      remaining_bandwidth = 10
    }]
  }]
  system_out_policy_map_name = "PM1"
  policy_map_statistics      = false
}
