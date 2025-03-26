resource "nxos_ospfv3_vrf" "example" {
  instance_name            = "OSPFv3"
  name                     = "VRF1"
  admin_state              = "enabled"
  bandwidth_reference      = 400000
  bandwidth_reference_unit = "mbps"
  router_id                = "34.56.78.90"
  bfd_control              = false
}
