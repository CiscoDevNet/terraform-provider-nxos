resource "nxos_physical_interface" "example" {
  interface_id             = "eth1/10"
  fec_mode                 = "auto"
  access_vlan              = "vlan-10"
  admin_state              = "up"
  auto_negotiation         = "on"
  bandwidth                = 1000
  delay                    = 10
  description              = "My Description"
  duplex                   = "auto"
  link_logging             = "enable"
  medium                   = "broadcast"
  mode                     = "trunk"
  mtu                      = 9216
  native_vlan              = "vlan-10"
  speed                    = "auto"
  speed_group              = "auto"
  trunk_vlans              = "100-200,300"
  uni_directional_ethernet = "disable"
}
