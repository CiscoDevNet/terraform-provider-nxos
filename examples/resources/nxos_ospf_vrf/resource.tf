resource "nxos_ospf_vrf" "example" {
  instance_name            = "OSPF1"
  name                     = "VRF1"
  admin_state              = "enabled"
  bandwidth_reference      = 400000
  bandwidth_reference_unit = "mbps"
  distance                 = 110
  router_id                = "34.56.78.90"
  control                  = "bfd"
}
