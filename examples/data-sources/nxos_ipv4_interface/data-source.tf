data "nxos_ipv4_interface" "example" {
  vrf = "default"
  interface_id = "eth1/10"
}
