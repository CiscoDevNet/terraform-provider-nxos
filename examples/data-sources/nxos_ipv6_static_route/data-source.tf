data "nxos_ipv6_static_route" "example" {
  vrf_name = "default"
  prefix   = "2001:db8:3333:4444:5555:6666:102:304/128"
}
