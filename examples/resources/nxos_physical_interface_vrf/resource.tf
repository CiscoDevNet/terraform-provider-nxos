resource "nxos_physical_interface_vrf" "example" {
  interface_id = "eth1/10"
  vrf_dn       = "sys/inst-default"
}
