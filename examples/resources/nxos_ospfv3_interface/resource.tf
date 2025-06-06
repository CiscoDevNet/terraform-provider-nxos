resource "nxos_ospfv3_interface" "example" {
  interface_id          = "eth1/10"
  advertise_secondaries = false
  area                  = "0.0.0.10"
  bfd                   = "disabled"
  cost                  = 1000
  dead_interval         = 60
  hello_interval        = 15
  network_type          = "p2p"
  passive               = "enabled"
  priority              = 10
}
