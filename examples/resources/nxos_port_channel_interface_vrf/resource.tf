resource "nxos_port_channel_interface_vrf" "example" {
  interface_id = "po1"
  vrf_dn       = "sys/inst-default"
}
