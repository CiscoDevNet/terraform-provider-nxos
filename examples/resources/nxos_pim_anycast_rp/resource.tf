resource "nxos_pim_anycast_rp" "example" {
  vrf_name         = "default"
  local_interface  = "eth1/10"
  source_interface = "eth1/10"
}
