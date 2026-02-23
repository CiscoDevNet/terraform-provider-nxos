resource "nxos_system" "example" {
  name = "LEAF1"
  default_qos_policy_interface_in = [{
    interface_id    = "eth1/10"
    policy_map_name = "PM1"
  }]
  policy_map_name      = "PM1"
  mtu                  = 9216
  default_admin_status = "up"
}
