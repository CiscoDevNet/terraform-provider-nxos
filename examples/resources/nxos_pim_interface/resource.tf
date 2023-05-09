resource "nxos_pim_interface" "example" {
  vrf_name = "default"
  interface_id = "eth1/10"
  admin_state = "enabled"
  bfd = "enabled"
  dr_priority = 10
  passive = false
  sparse_mode = true
}
