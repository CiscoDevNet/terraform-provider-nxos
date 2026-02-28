resource "nxos_hsrp" "example" {
  admin_state                          = "enabled"
  instance_admin_state                 = "enabled"
  bfd                                  = "enabled"
  control                              = "stateful-ha"
  extended_hold_interval               = 30
  extended_hold_interval_configuration = "disabled"
  interfaces = [{
    interface_id                       = "vlan10"
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
    groups = [{
      group_id                              = 1
      address_family                        = "ipv4"
      authentication_md5_compatibility_mode = "disabled"
      authentication_md5_key_chain_name     = "myKeyChain"
      authentication_md5_key_name           = "myKey"
      authentication_md5_key_string_type    = "unencrypted"
      authentication_md5_timeout            = 100
      authentication_md5_type               = "key-string"
      authentication_secret                 = "cisco123"
      authentication_type                   = "simple"
      control                               = "preempt"
      follow                                = "masterGroup"
      forwarding_lower_threshold            = 1
      hello_interval                        = 1000
      hold_interval                         = 3000
      ip_address                            = "10.0.0.1"
      ip_obtain_mode                        = "admin"
      mac_address                           = "00:00:00:00:00:01"
      name                                  = "myHsrpGroup"
      preempt_delay_minimum                 = 10
      preempt_delay_reload                  = 10
      preempt_delay_sync                    = 10
      priority                              = 110
    }]
  }]
}
