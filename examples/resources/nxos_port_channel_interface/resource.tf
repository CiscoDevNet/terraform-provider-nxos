resource "nxos_port_channel_interface" "example" {
  interface_id          = "po1"
  port_channel_mode     = "active"
  minimum_links         = 2
  maximum_links         = 10
  suspend_individual    = "disable"
  access_vlan           = "unknown"
  admin_state           = "up"
  auto_negotiation      = "on"
  bandwidth             = 1000
  delay                 = 10
  description           = "My Description"
  duplex                = "auto"
  layer                 = "Layer3"
  link_logging          = "enable"
  medium                = "broadcast"
  mode                  = "access"
  mtu                   = 1500
  native_vlan           = "unknown"
  speed                 = "auto"
  trunk_vlans           = "1-4094"
  user_configured_flags = "admin_layer,admin_mtu,admin_state"
}
