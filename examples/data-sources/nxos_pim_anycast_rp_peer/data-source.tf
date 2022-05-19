data "nxos_pim_anycast_rp_peer" "example" {
  vrf_name       = "default"
  address        = "10.1.1.1/32"
  rp_set_address = "20.1.1.1/32"
}
