resource "nxos_pim_static_rp_policy" "example" {
  vrf_name = "default"
  name = "RP"
}
