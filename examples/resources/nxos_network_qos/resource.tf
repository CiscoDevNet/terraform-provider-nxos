resource "nxos_network_qos" "example" {
  class_maps = {
    "Voice" = {
      match_type = "match-any"
      cos_values = {
        "4" = {}
      }
    }
  }
  policy_maps = {
    "PM1" = {
      match_type = "match-any"
      match_class_maps = {
        "Voice" = {
          mtu_value = 9216
        }
      }
    }
  }
  system_in_policy_map_name = "PM1"
  policy_map_statistics     = false
}
