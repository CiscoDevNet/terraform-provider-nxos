resource "nxos_port_channel_interface_member" "example" {
  interface_id = "po1"
  interface_dn = "sys/intf/phys-[eth1/11]"
}
