data "nxos_ipv4_interface_address" "example" {
  vrf          = "default"
  interface_id = "eth1/59"
  address      = "1.1.1.1/24"
}
