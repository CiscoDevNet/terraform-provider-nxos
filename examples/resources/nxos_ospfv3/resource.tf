resource "nxos_ospfv3" "example" {
  admin_state = "enabled"
  instances = [{
    name         = "OSPFv3"
    admin_state  = "enabled"
    flush_routes = false
    isolate      = false
    vrfs = [{
      name                      = "VRF1"
      admin_state               = "enabled"
      bandwidth_reference       = 400000
      bandwidth_reference_unit  = "mbps"
      router_id                 = "34.56.78.90"
      bfd_control               = false
      log_adjacency_changes     = "brief"
      discard_route_external    = false
      discard_route_internal    = false
      name_lookup               = true
      passive_interface_default = false
      areas = [{
        area_id                  = "0.0.0.10"
        redistribute             = false
        nssa_translator_role     = "always"
        summary                  = false
        suppress_forward_address = false
        type                     = "regular"
      }]
      address_families = [{
        address_family_type           = "ipv6-ucast"
        administrative_distance       = "10"
        default_metric                = "1024"
        default_route_nssa_pbit_clear = true
        max_ecmp_cost                 = 16
      }]
    }]
  }]
  interfaces = [{
    interface_id          = "eth1/4"
    advertise_secondaries = false
    area                  = "0.0.0.10"
    bfd_control           = "disabled"
    cost                  = 1000
    dead_interval         = 60
    hello_interval        = 15
    network_type          = "p2p"
    passive               = "enabled"
    priority              = 10
    admin_state           = "enabled"
    instance              = "OSPFv3"
    instance_id           = 1
    mtu_ignore            = true
    retransmit_interval   = 10
    transmit_delay        = 5
  }]
}
