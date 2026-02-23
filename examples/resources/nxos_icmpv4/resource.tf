resource "nxos_icmpv4" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  vrfs = [{
    name = "VRF1"
    interfaces = [{
      id      = "vlan10"
      control = "port-unreachable"
    }]
  }]
}
