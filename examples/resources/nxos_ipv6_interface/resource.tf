resource "nxos_ipv6_interface" "example" {
  vrf                    = "default"
  interface_id           = "eth1/10"
  auto_configuration     = "disabled"
  default_route          = "disabled"
  ipv6_forward           = "disabled"
  forward                = "disabled"
  link_address_use_bia   = "disabled"
  use_link_local_address = "disabled"
  urpf                   = "disabled"
  link_local_address     = "2001:db8:3333:4444:5555:6666:7777:8888"
}
