data "nxos_ospfv3_area_address_family" "example" {
  instance_name = "OSPFv3"
  vrf_name      = "default"
  area_id       = "0.0.0.10"
  family_type   = "ipv6-ucast"
}
