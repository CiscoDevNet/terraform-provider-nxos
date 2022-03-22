data "nxos_ipv4_interface_address" "example" {
  vrf          = "default"
  interface_id = "eth1/59"
  address      = "24.63.46.49/30"
}
