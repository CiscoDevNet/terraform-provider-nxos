data "nxos_ospf_interface" "example" {
  instance_name = "OSPF1"
  vrf_name      = "default"
  interface_id  = "eth1/10"
}
