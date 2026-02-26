resource "nxos_default_qos" "example" {
  class_maps = [{
    name       = "Voice"
    match_type = "match-any"
    dscp_values = [{
      value = "ef"
    }]
  }]
  policy_maps = [{
    name       = "PM1"
    match_type = "match-any"
    match_class_maps = [{
      name                          = "Voice"
      set_qos_group_id              = 1
      police_bc_rate                = 200
      police_bc_unit                = "mbytes"
      police_be_rate                = 200
      police_be_unit                = "mbytes"
      police_cir_rate               = 10000
      police_cir_unit               = "mbps"
      police_conform_action         = "transmit"
      police_conform_set_cos        = 0
      police_conform_set_dscp       = 0
      police_conform_set_precedence = "routine"
      police_conform_set_qos_group  = 0
      police_exceed_action          = "transmit"
      police_exceed_set_cos         = 0
      police_exceed_set_dscp        = 0
      police_exceed_set_precedence  = "routine"
      police_exceed_set_qos_group   = 0
      police_pir_rate               = 10000
      police_pir_unit               = "mbps"
      police_violate_action         = "drop"
      police_violate_set_cos        = 0
      police_violate_set_dscp       = 0
      police_violate_set_precedence = "routine"
      police_violate_set_qos_group  = 0
    }]
  }]
  policy_interface_in = [{
    interface_id    = "eth1/10"
    policy_map_name = "PM1"
  }]
}
