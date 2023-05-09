data "nxos_pim_interface" "example" {
  vrf_name = "default"
  interface_id = "eth1/10"
}
