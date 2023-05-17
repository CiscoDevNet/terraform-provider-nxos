resource "nxos_ipv4_access_list_policy_egress_interface_instance" "example" {
  interface_id = "eth1/10"
  name         = "ACL1"
}
