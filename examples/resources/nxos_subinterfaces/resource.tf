resource "nxos_subinterfaces" "example" {
  items = [{
    interface_id = "eth1/10.124"
    admin_state  = "down"
    bandwidth    = 1000
    delay        = 10
    description  = "My Description"
    encap        = "vlan-124"
    link_logging = "enable"
    medium       = "broadcast"
    mtu          = 1500
  }]
}
