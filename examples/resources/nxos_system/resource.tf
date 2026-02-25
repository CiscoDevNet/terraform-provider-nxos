resource "nxos_system" "example" {
  name                 = "LEAF1"
  policy_map_name      = "PM1"
  mtu                  = 9216
  default_admin_status = "up"
}
