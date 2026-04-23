resource "nxos_hsrp" "example" {
  admin_state                          = "enabled"
  instance_admin_state                 = "enabled"
  bfd                                  = "enabled"
  control                              = "stateful-ha"
  extended_hold_interval               = 30
  extended_hold_interval_configuration = "disabled"
  interfaces = {
    "vlan10" = {
      admin_state                        = "enabled"
      bfd                                = "disabled"
      bia_scope                          = "local"
      control                            = "bia"
      delay_minimum                      = 5
      description                        = "My HSRP interface"
      mac_refresh_interval               = 30
      mac_refresh_interval_configuration = "disabled"
      name                               = "myHsrpIf"
      reload_delay                       = 5
      version                            = "v2"
      groups = {
        "1;ipv4" = {
          authentication_type        = "simple"
          control                    = "preempt"
          forwarding_lower_threshold = 1
          hello_interval             = 1000
          hold_interval              = 3000
          ip_address                 = "10.0.0.1"
          ip_obtain_mode             = "admin"
          name                       = "myHsrpGroup"
          preempt_delay_minimum      = 10
          preempt_delay_reload       = 10
          preempt_delay_sync         = 10
          priority                   = 110
        }
      }
    }
  }
}
