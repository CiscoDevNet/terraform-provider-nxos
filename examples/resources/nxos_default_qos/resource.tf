resource "nxos_default_qos" "example" {
  class_maps = {
    "Voice" = {
      match_type = "match-any"
      dscp_values = {
        "ef" = {}
      }
      not_dscp_values = {
        "af11" = {}
      }
      cos_values = {
        "3" = {}
      }
      not_cos_values = {
        "5" = {}
      }
      precedence_values = {
        "critical" = {}
      }
      not_precedence_values = {
        "flash" = {}
      }
      access_groups = {
        "ACL1" = {}
      }
    }
  }
  policy_maps = {
    "PM1" = {
      match_type = "match-any"
      match_class_maps = {
        "Voice" = {
          set_qos_group_id              = 1
          set_dscp_tunnel               = false
          set_dscp_value                = "af21"
          set_cos_value                 = "3"
          set_precedence_tunnel         = false
          set_precedence_value          = "critical"
          police_bc_rate                = 200
          police_bc_unit                = "mbytes"
          police_be_rate                = 200
          police_be_unit                = "mbytes"
          police_cir_rate               = 10000
          police_cir_unit               = "mbps"
          police_conform_action         = "set-dscp-transmit"
          police_conform_set_cos        = 0
          police_conform_set_dscp       = "af22"
          police_conform_set_precedence = "routine"
          police_conform_set_qos_group  = 0
          police_exceed_action          = "set-dscp-transmit"
          police_exceed_set_cos         = 0
          police_exceed_set_dscp        = "af31"
          police_exceed_set_precedence  = "routine"
          police_exceed_set_qos_group   = 0
          police_pir_rate               = 10000
          police_pir_unit               = "mbps"
          police_violate_action         = "set-dscp-transmit"
          police_violate_set_cos        = 0
          police_violate_set_dscp       = "cs1"
          police_violate_set_precedence = "routine"
          police_violate_set_qos_group  = 0
        }
      }
    }
  }
  policy_interface_in = {
    "eth1/10" = {
      policy_map_name       = "PM1"
      policy_map_statistics = false
    }
  }
  policy_nve_interface_in = {
    "1" = {
      policy_map_name       = "PM1"
      policy_map_statistics = false
    }
  }
}
