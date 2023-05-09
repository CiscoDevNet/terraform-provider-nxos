data "nxos_pim_static_rp_group_list" "example" {
  vrf_name = "default"
  rp_address = "1.2.3.4"
  address = "224.0.0.0/4"
}
