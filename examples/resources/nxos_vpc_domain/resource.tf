resource "nxos_vpc_domain" "example" {
  admin_state                    = "enabled"
  domain_id                      = 100
  auto_recovery                  = "enabled"
  auto_recovery_interval         = 480
  delay_restore_orphan_port      = 10
  delay_restore_svi              = 20
  delay_restore_vpc              = 60
  dscp                           = 0
  fast_convergence               = "enabled"
  graceful_consistency_check     = "disabled"
  l3_peer_router                 = "enabled"
  l3_peer_router_syslog          = "enabled"
  l3_peer_router_syslog_interval = 3000
  peer_gateway                   = "enabled"
  peer_ip                        = "1.2.3.4"
  peer_switch                    = "enabled"
  role_priority                  = 100
  sys_mac                        = "01:01:01:01:01:01"
  system_priority                = 100
  track                          = 10
  virtual_ip                     = "1.1.1.1"
}
