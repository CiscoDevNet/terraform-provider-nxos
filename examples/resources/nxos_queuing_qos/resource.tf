resource "nxos_queuing_qos" "example" {
  policy_maps = {
    "PM1" = {
      match_type = "match-any"
      match_class_maps = {
        "c-out-q1" = {
          priority           = 3
          cap_average        = false
          drop_availability  = 7
          random_detect_ecn  = true
          max_threshold      = 3000
          max_threshold_unit = "kbytes"
          min_threshold      = 150
          min_threshold_unit = "kbytes"
          optimization       = "mesh"
          weight             = 5
        }
      }
    }
  }
  system_out_policy_map_name = "PM1"
  policy_map_statistics      = false
}
