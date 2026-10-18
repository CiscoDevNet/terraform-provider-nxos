resource "nxos_queuing_qos" "example" {
  policy_maps = {
    "PM1" = {
      match_type = "match-any"
      match_class_maps = {
        "c-out-q1" = {
          priority            = 3
          cap_average         = false
          drop_availability   = 7
          random_detect_ecn   = true
          max_threshold       = 3000
          max_threshold_unit  = "kbytes"
          min_threshold       = 150
          min_threshold_unit  = "kbytes"
          optimization        = "mesh"
          weight              = 5
          shape_max_rate      = 2000
          shape_max_rate_unit = "mbps"
          shape_min_rate      = 1000
          shape_min_rate_unit = "mbps"
          queue_limit_unit    = "kbytes"
          queue_limit         = 1000
          burst_detect_state  = "enabled"
        }
        "c-out-q2" = {
          remaining_bandwidth = 10
        }
        "c-out-q3" = {
          remaining_bandwidth = 10
        }
        "c-out-q-default" = {
          remaining_bandwidth = 10
        }
      }
    }
  }
  system_out_policy_map_name = "PM1"
  policy_map_statistics      = false
  policy_interface_out = {
    "eth1/10" = {
      policy_map_name       = "PM1"
      policy_map_statistics = false
    }
  }
}
