resource "nxos_ipv6_interface_address" "example" {
  vrf          = "default"
  interface_id = "eth1/10"
  address      = "2001:db8:3333:4444:5555:6666:7777:8888"
  type         = "primary"
  tag          = 1234
}
