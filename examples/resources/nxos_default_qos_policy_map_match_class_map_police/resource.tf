resource "nxos_default_qos_policy_map_match_class_map_police" "example" {
  policy_map_name        = "PM1"
  class_map_name         = "Voice"
  bc_rate                = 200
  bc_unit                = "mbytes"
  be_rate                = 200
  be_unit                = "mbytes"
  cir_rate               = 10000
  cir_unit               = "mbps"
  conform_action         = "transmit"
  conform_set_cos        = 0
  conform_set_dscp       = 0
  conform_set_precedence = "routine"
  conform_set_qos_group  = 0
  exceed_action          = "transmit"
  exceed_set_cos         = 0
  exceed_set_dscp        = 0
  exceed_set_precedence  = "routine"
  exceed_set_qos_group   = 0
  pir_rate               = 10000
  pir_unit               = "mbps"
  violate_action         = "drop"
  violate_set_cos        = 0
  violate_set_dscp       = 0
  violate_set_precedence = "routine"
  violate_set_qos_group  = 0
}
