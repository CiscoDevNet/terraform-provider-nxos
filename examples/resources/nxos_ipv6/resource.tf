resource "nxos_ipv6" "example" {
  admin_state                             = "enabled"
  instance_access_list_match_local        = "enabled"
  instance_admin_state                    = "enabled"
  instance_control                        = "stateful-ha"
  instance_drop_nd_fragments              = "enabled"
  instance_queue_packets                  = "enabled"
  instance_static_neighbor_outside_subnet = "enabled"
  instance_switch_packets                 = "all"
  vrfs = [{
    name = "VRF1"
    static_routes = [{
      prefix      = "2001:db8:3333:4444:5555:6666:102:304/128"
      control     = "bfd"
      description = "My Description"
      preference  = 10
      tag         = 100
      next_hops = [{
        interface_id          = "unspecified"
        address               = "a:b::c:d/128"
        vrf_name              = "default"
        description           = "My Description"
        object                = 10
        preference            = 123
        tag                   = 10
        route_name            = "nh-name"
        rewrite_encapsulation = "vlan-1"
      }]
    }]
    interfaces = [{
      interface_id               = "eth1/10"
      auto_configuration         = "disabled"
      default_route              = "disabled"
      forward                    = "disabled"
      link_local_address_use_bia = "disabled"
      use_link_local_address     = "disabled"
      urpf                       = "disabled"
      link_local_address         = "2001:db8:3333:4444:5555:6666:7777:8888"
      addresses = [{
        address                 = "2001:db8:3333:4444:5555:6666:7777:8888"
        type                    = "primary"
        tag                     = 1234
        aggregate_prefix_length = 64
        control                 = "anycast"
        preference              = 10
        use_bia                 = "enabled"
        vpc_peer                = "2001:db8::1"
      }]
    }]
  }]
}
