resource "nxos_icmp_interface" "example" {
  domain_name  = "icmpv4dom1"
  interface_id = "vlan10"
  control      = ""
}
