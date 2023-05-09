resource "nxos_hmm_interface" "example" {
  interface_id = "vlan10"
  admin_state = "enabled"
  mode = "anycastGW"
}
