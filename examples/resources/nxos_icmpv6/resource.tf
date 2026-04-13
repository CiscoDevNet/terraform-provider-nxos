resource "nxos_icmpv6" "example" {
  admin_state                = "enabled"
  adjacency_stale_timer      = 500
  adjacency_stale_timer_icmp = "disabled"
  instance_admin_state       = "enabled"
  control                    = "stateful-ha"
  redirect_syslog            = "disabled"
  redirect_syslog_interval   = 120
  interfaces = {
    "vlan10" = {
      control = "unreachables"
    }
  }
}
