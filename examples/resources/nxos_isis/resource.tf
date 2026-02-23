resource "nxos_isis" "example" {
  admin_state = "enabled"
  instances = [{
    name        = "ISIS1"
    admin_state = "enabled"
    vrfs = [{
      name                    = "default"
      admin_state             = "enabled"
      authentication_check_l1 = false
      authentication_check_l2 = false
      authentication_key_l1   = ""
      authentication_key_l2   = ""
      authentication_type_l1  = "unknown"
      authentication_type_l2  = "unknown"
      bandwidth_reference     = 400000
      banwidth_reference_unit = "mbps"
      is_type                 = "l2"
      metric_type             = "wide"
      mtu                     = 2000
      net                     = "49.0001.0000.0000.3333.00"
      passive_default         = "l12"
      address_families = [{
        address_family              = "v4"
        segment_routing_mpls        = true
        enable_bfd                  = false
        prefix_advertise_passive_l1 = true
        prefix_advertise_passive_l2 = true
      }]
      overload_startup_time = 60
    }]
  }]
  interfaces = [{
    interface_id            = "eth1/10"
    authentication_check    = false
    authentication_check_l1 = false
    authentication_check_l2 = false
    authentication_key      = ""
    authentication_key_l1   = ""
    authentication_key_l2   = ""
    authentication_type     = "unknown"
    authentication_type_l1  = "unknown"
    authentication_type_l2  = "unknown"
    circuit_type            = "l2"
    vrf                     = "default"
    hello_interval          = 20
    hello_interval_l1       = 20
    hello_interval_l2       = 20
    hello_multiplier        = 4
    hello_multiplier_l1     = 4
    hello_multiplier_l2     = 4
    hello_padding           = "never"
    instance_name           = "ISIS1"
    metric_l1               = 1000
    metric_l2               = 1000
    mtu_check               = true
    mtu_check_l1            = true
    mtu_check_l2            = true
    network_type_p2p        = "on"
    passive                 = "l1"
    priority_l1             = 80
    priority_l2             = 80
    enable_ipv4             = true
  }]
}
