data "nxos_ipv4_static_route" "example" {
  vrf_name = "default"
  prefix   = "1.1.1.0/24"
}
