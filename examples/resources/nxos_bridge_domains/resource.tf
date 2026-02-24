resource "nxos_bridge_domains" "example" {
  items = [{
    fabric_encap = "vlan-10"
    access_encap = "unknown"
    name         = "VLAN10"
  }]
}
