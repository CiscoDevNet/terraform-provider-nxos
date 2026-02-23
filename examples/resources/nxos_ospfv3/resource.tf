resource "nxos_ospfv3" "example" {
  admin_state = "enabled"
  instances = [{
    name        = "OSPFv3"
    admin_state = "enabled"
    vrfs = [{
      name                     = "VRF1"
      admin_state              = "enabled"
      bandwidth_reference      = 400000
      bandwidth_reference_unit = "mbps"
      router_id                = "34.56.78.90"
      bfd_control              = false
      areas = [{
        area_id                  = "0.0.0.10"
        redistribute             = false
        summary                  = false
        suppress_forward_address = false
        type                     = "regular"
      }]
      address_families = [{
        address_family_type     = "ipv6-ucast"
        administrative_distance = "10"
        default_metric          = "1024"
        max_ecmp_cost           = 16
      }]
    }]
  }]
  interfaces = [{
    interface_id          = "eth1/10"
    advertise_secondaries = false
    area                  = "0.0.0.10"
    bfd                   = "disabled"
    cost                  = 1000
    dead_interval         = 60
    hello_interval        = 15
    network_type          = "p2p"
    passive               = "enabled"
    priority              = 10
  }]
}
