resource "nxos_default_qos_class_map" "example" {
  name       = "Voice"
  match_type = "match-any"
  dscp_values = [{
    value = "ef"
  }]
}
