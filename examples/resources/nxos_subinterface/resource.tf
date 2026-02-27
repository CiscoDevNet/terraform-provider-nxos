resource "nxos_subinterface" "example" {
  interface_id            = "eth1/10.124"
  admin_state             = "down"
  bandwidth               = 1000
  delay                   = 10
  description             = "My Description"
  encap                   = "vlan-124"
  link_logging            = "enable"
  medium                  = "broadcast"
  mtu                     = 1500
  mtu_inherit             = false
  router_mac              = "AA:BB:CC:DD:EE:FF"
  router_mac_ipv6_extract = "disable"
  snmp_trap               = "disable"
  vrf_dn                  = "sys/inst-VRF123"
}
