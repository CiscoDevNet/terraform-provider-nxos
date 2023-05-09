resource "nxos_ipv4_access_list_policy_ingress_interface_instace" "example" {
  interface_id = "eth1/10"
  name         = "ACL1"
}
