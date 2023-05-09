resource "nxos_bridge_domain" "example" {
  fabric_encap = "vlan-10"
  access_encap = "unknown"
  name         = "VLAN10"
}
