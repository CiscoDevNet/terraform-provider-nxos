resource "nxos_default_qos_class_map_dscp" "example" {
  class_map_name = "Voice"
  value          = "ef"
}
