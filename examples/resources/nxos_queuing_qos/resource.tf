resource "nxos_queuing_qos" "example" {
  policy_maps = {
    "PM1" = {
      match_type = "match-any"
      match_class_maps = {
        "c-out-q1" = {
          priority = 1
        }
      }
    }
  }
  system_out_policy_map_name = "PM1"
  policy_map_statistics      = false
}
