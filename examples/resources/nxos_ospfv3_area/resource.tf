resource "nxos_ospfv3_area" "example" {
  instance_name          = "OSPFv3"
  vrf_name               = "VRF1"
  area_id                = "0.0.0.10"
  redistribute           = false
  summary                = false
  supress_foward_address = false
  type                   = "regular"
}
