resource "nxos_dhcp_relay_address" "example" {
  interface_id = "eth1/10"
  vrf          = "VRF1"
  address      = "1.1.1.1"
}
