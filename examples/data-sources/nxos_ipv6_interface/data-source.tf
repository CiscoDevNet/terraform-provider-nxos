data "nxos_ipv6_interface" "example" {
  vrf          = "default"
  interface_id = "eth1/10"
}
