resource "nxos_default_qos_policy_interface_in_policy_map" "example" {
  interface_id    = "eth1/10"
  policy_map_name = "PM1"
}
