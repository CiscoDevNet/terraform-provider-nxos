resource "nxos_pim_static_rp" "example" {
  vrf_name = "default"
  address  = "1.2.3.4"
}
