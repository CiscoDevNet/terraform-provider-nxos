resource "nxos_loopback_interface" "example" {
  interface_id = "lo123"
  admin_state  = "down"
  description  = "My Description"
  link_logging = "enable"
  vrf_dn       = "sys/inst-default"
}
