resource "nxos_physical_interface" "example" {
  interface_id             = "eth1/10"
  fec_mode                 = "auto"
  access_vlan              = "unknown"
  admin_state              = "up"
  auto_negotiation         = "on"
  bandwidth                = 1000
  delay                    = 10
  description              = "My Description"
  duplex                   = "auto"
  layer                    = "Layer3"
  link_logging             = "enable"
  medium                   = "broadcast"
  mode                     = "access"
  mtu                      = 1500
  native_vlan              = "unknown"
  speed                    = "auto"
  speed_group              = "auto"
  trunk_vlans              = "1-4094"
  uni_directional_ethernet = "disable"
}
