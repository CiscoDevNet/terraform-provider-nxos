resource "nxos_hmm" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  anycast_mac          = "20:20:00:00:10:10"
  interfaces = [{
    interface_id = "vlan10"
    admin_state  = "enabled"
    mode         = "anycastGW"
  }]
}
