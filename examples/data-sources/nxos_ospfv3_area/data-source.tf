data "nxos_ospfv3_area" "example" {
  instance_name = "nac-ospfv3"
  vrf_name      = "VRF1"
  area_id       = "0.0.0.10"
}
