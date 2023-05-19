resource "nxos_ipv4_access_list_policy_egress_interface" "example" {
  interface_id     = "eth1/10"
  access_list_name = "ACL1"
}
