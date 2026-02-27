resource "nxos_system" "example" {
  name                                      = "LEAF1"
  mtu                                       = 9216
  default_admin_state                       = "up"
  admin_link_down_syslog_level              = 4
  admin_link_up_syslog_level                = 4
  admin_state                               = "enabled"
  allow_unsupported_sfp                     = true
  chassis_infrastructure_adaptor_vlan       = 100
  chassis_infrastructure_epds_port_number   = 100
  chassis_infrastructure_ipv6_address       = "2001:db8::1"
  chassis_infrastructure_vlan               = 100
  chassis_management_instance               = "mgmt0"
  chassis_management_instance_fabric_number = "LeftFabric"
  control                                   = "stateful-ha"
  interface_syslog_info                     = "info-1"
  log_event                                 = "linkStatusEnable"
  default_layer                             = "Layer3"
  system_interface_admin_state              = "up"
  system_link_failure_laser_on              = false
  system_storm_control_multi_threshold      = false
  vlan_tag_native                           = false
}
