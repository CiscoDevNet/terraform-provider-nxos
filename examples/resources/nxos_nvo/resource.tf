resource "nxos_nvo" "example" {
  nve_interfaces = [{
    id                               = 1
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
    vnis = [{
      vni                           = 103100
      associate_vrf                 = false
      multicast_group               = "0.0.0.0"
      multisite_ingress_replication = "disable"
      suppress_arp                  = "off"
      protocol                      = "bgp"
    }]
  }]
}
