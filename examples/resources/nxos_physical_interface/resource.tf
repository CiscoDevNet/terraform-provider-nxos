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
  link_debounce_down       = 200
  link_debounce_up         = 0
  medium                   = "broadcast"
  mode                     = "access"
  mtu                      = 1500
  native_vlan              = "unknown"
  speed                    = "auto"
  speed_group              = "auto"
  trunk_vlans              = "1-4094"
  uni_directional_ethernet = "disable"
  user_configured_flags    = "admin_layer,admin_mtu,admin_state"
}
