resource "nxos_poe" "example" {
  instance_admin_state = "enabled"
  instance_control     = "stateful-ha"
  logging_level        = 3
  interfaces = {
    "eth1/10" = {
      max_consumption          = 20000
      port_default_consumption = 10000
      port_mode                = "static"
      port_police              = "log"
      port_priority            = "high"
    }
  }
}
