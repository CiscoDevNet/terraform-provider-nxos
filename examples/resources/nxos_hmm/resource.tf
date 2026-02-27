resource "nxos_hmm" "example" {
  admin_state             = "enabled"
  instance_admin_state    = "enabled"
  anycast_mac             = "20:20:00:00:10:10"
  administrative_distance = 150
  control                 = "stateful-ha"
  limit_vlan_mac          = 100
  selective_host_probe    = "yes"
  interfaces = [{
    interface_id = "vlan10"
    admin_state  = "enabled"
    mode         = "anycastGW"
    description  = "My Description"
  }]
}
