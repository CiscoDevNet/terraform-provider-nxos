resource "nxos_nve_interface" "example" {
  admin_state                      = "enabled"
  advertise_virtual_mac            = true
  hold_down_time                   = 60
  host_reachability_protocol       = "bgp"
  ingress_replication_protocol_bgp = true
  multisite_source_interface       = "unspecified"
  source_interface                 = "lo0"
  supress_mac_route                = true
}
