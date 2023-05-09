resource "nxos_pim_vrf" "example" {
  name = "default"
  admin_state = "enabled"
  bfd = true
}
