resource "nxos_svi_interface_vrf" "example" {
  interface_id = "vlan293"
  vrf_dn = "sys/inst-VRF123"
}
