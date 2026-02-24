resource "nxos_ipv6_vrfs" "example" {
  items = [{
    name = "VRF1"
  }]
}
