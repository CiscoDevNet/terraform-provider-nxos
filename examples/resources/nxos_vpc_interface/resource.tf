resource "nxos_vpc_interface" "example" {
  vpc_interface_id          = 1
  port_channel_interface_dn = "sys/intf/aggr-[po1]"
}
