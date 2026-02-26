resource "nxos_system" "example" {
  name                = "LEAF1"
  mtu                 = 9216
  default_admin_state = "up"
}
