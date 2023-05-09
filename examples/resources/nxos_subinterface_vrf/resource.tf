resource "nxos_subinterface_vrf" "example" {
  interface_id = "eth1/10.124"
  vrf_dn = "sys/inst-VRF123"
}
