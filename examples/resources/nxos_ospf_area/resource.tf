resource "nxos_ospf_area" "example" {
  instance_name       = "OSPF1"
  vrf_name            = "VRF1"
  area_id             = "0.0.0.10"
  authentication_type = "none"
  cost                = 10
  type                = "stub"
}
