resource "nxos_vrf" "example" {
  name        = "VRF1"
  description = "My VRF1 Description"
  encap       = "vxlan-103901"
}
