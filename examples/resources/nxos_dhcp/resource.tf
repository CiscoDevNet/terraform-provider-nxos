resource "nxos_dhcp" "example" {
  relay_interfaces = [{
    interface_id = "eth1/10"
    addresses = [{
      vrf     = "VRF1"
      address = "1.1.1.1"
    }]
  }]
}
