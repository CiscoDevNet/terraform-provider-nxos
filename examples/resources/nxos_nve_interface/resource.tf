resource "nxos_nve_interface" "example" {
  admin_state                      = "enabled"
  advertise_virtual_mac            = true
  hold_down_time                   = 60
  host_reachability_protocol       = "bgp"
  ingress_replication_protocol_bgp = true
  multicast_group_l2               = "0.0.0.0"
  multicast_group_l3               = "0.0.0.0"
  multisite_source_interface       = "unspecified"
  source_interface                 = "lo0"
  suppress_arp                     = true
  suppress_mac_route               = false
}
