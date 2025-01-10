resource "nxos_ipv4_interface_address" "example" {
  vrf          = "default"
  interface_id = "eth1/10"
  address      = "24.63.46.49/30"
  type         = "primary"
  tag          = 1234
}
