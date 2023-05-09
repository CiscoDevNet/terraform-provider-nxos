data "nxos_ospf_interface" "example" {
  instance_name = "OSPF1"
  vrf_name = "VRF1"
  interface_id = "eth1/10"
}
