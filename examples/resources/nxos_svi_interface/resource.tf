resource "nxos_svi_interface" "example" {
  interface_id            = "vlan293"
  admin_state             = "down"
  autostate               = false
  bandwidth               = 1000
  carrier_delay           = 200
  delay                   = 10
  description             = "My Description"
  inband_management       = false
  load_interval_counter_1 = 90
  load_interval_counter_2 = 120
  load_interval_counter_3 = 90
  mac_address             = "00:25:B5:00:00:01"
  medium                  = "bcast"
  mtu                     = 9216
  mtu_inherit             = false
  snmp_trap_link_status   = false
  vlan_id                 = 100
  vrf_dn                  = "sys/inst-VRF123"
}
