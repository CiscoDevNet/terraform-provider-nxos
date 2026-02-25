resource "nxos_bridge_domains" "example" {
  bridge_domains = [{
    fabric_encap = "vlan-10"
    access_encap = "unknown"
    name         = "VLAN10"
  }]
}
