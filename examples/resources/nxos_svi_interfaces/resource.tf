resource "nxos_svi_interfaces" "example" {
  items = [{
    interface_id = "vlan293"
    admin_state  = "down"
    bandwidth    = 1000
    delay        = 10
    description  = "My Description"
    medium       = "bcast"
    mtu          = 9216
  }]
}
