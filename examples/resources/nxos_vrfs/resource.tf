resource "nxos_vrfs" "example" {
  items = [{
    name        = "VRF1"
    description = "My VRF1 Description"
    encap       = "vxlan-103901"
  }]
}
