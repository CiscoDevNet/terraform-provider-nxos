resource "nxos_pim_ssm_range" "example" {
  vrf_name     = "default"
  group_list_1 = "232.0.0.0/8"
  group_list_2 = "233.0.0.0/8"
  group_list_3 = "0.0.0.0"
  group_list_4 = "0.0.0.0"
  prefix_list  = ""
  route_map    = ""
  ssm_none     = false
}
