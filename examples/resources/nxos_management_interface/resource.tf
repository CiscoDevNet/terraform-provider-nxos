resource "nxos_management_interface" "example" {
  management_interfaces = {
    "mgmt0" = {
      admin_state      = "up"
      auto_negotiation = "on"
      description      = "My Description"
      duplex           = "auto"
      mtu              = 1500
      snmp_trap_state  = "enable"
      speed            = "auto"
      vrf_dn           = "sys/inst-management"
    }
  }
}
