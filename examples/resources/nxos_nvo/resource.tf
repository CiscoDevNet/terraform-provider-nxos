resource "nxos_nvo" "example" {
  vxlan_udp_port             = 4789
  vxlan_udp_source_port_mode = "high"
  nve_interfaces = [{
    id                                 = 1
    admin_state                        = "enabled"
    advertise_virtual_mac              = true
    hold_down_time                     = 60
    host_reachability_protocol         = "bgp"
    ingress_replication_protocol_bgp   = true
    multicast_group_l2                 = "0.0.0.0"
    multicast_group_l3                 = "0.0.0.0"
    multisite_source_interface         = "unspecified"
    source_interface                   = "lo0"
    suppress_arp                       = true
    suppress_mac_route                 = false
    anycast_source_interface           = "unspecified"
    configuration_source               = "cli"
    controller_id                      = 0
    description                        = "My NVE interface"
    encapsulation_type                 = "vxlan"
    fabric_ready_time                  = 30
    multicast_routing_source_interface = "unspecified"
    multisite_virtual_mac              = "00:00:00:00:00:00"
    suppress_nd                        = false
    virtual_mac                        = "00:00:00:00:00:00"
    vnis = [{
      vni                           = 103100
      associate_vrf                 = false
      multicast_group               = "0.0.0.0"
      multisite_ingress_replication = "disable"
      suppress_arp                  = "off"
      legacy_mode                   = false
      multisite_multicast_group     = "0.0.0.0"
      spine_anycast_gateway         = false
      ingress_replication_protocol  = "bgp"
    }]
  }]
}
