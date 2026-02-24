resource "nxos_loopback_interfaces" "example" {
  items = [{
    interface_id = "lo123"
    admin_state  = "down"
    description  = "My Description"
  }]
}
