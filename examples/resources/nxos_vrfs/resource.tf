resource "nxos_vrfs" "example" {
  items = [{
    name          = "VRF1"
    description   = "My VRF1 Description"
    admin_state   = "admin-up"
    controller_id = 1
    encap         = "vxlan-103901"
    l3vni         = false
    oui           = "000000"
    vpn_id        = "100:200"
  }]
}
