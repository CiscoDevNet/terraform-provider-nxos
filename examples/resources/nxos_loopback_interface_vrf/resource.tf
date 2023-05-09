resource "nxos_loopback_interface_vrf" "example" {
  interface_id = "lo123"
  vrf_dn       = "sys/inst-default"
}
