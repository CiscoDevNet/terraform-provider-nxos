resource "nxos_ipv4_vrfs" "example" {
  items = [{
    name = "VRF1"
  }]
}
