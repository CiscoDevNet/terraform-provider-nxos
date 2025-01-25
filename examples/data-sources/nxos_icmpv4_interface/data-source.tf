data "nxos_icmpv4_interface" "example" {
  vrf_name     = "VRF1"
  interface_id = "vlan10"
}
