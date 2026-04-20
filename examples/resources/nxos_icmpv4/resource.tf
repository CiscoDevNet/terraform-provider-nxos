resource "nxos_icmpv4" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  control              = "stateful-ha"
  vrfs = {
    "VRF1" = {
      interfaces = {
        "vlan10" = {
          control = "port-unreachable"
        }
      }
    }
  }
}
