resource "nxos_igmp_snooping" "example" {
  admin_state                                = "enabled"
  instance_admin_state                       = "enabled"
  instance_control                           = "stateful-ha"
  domain_control                             = "opt-flood"
  domain_name                                = "default"
  global_vlan_vxlan                          = true
  global_vlan_disable_nve_static_router_port = false
  global_vlan_vxlan_umc_drop_vlan            = "100-200"
}
