resource "nxos_ipv4_interface" "example" {
  vrf          = "default"
  interface_id = "eth1/10"
  drop_glean   = "disabled"
  forward      = "disabled"
  unnumbered   = "unspecified"
  urpf         = "disabled"
}
