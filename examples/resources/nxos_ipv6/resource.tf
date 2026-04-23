resource "nxos_ipv6" "example" {
  access_list_match_local        = "enabled"
  admin_state                    = "enabled"
  control                        = "stateful-ha"
  drop_nd_fragments              = "enabled"
  queue_packets                  = "enabled"
  static_neighbor_outside_subnet = "enabled"
  switch_packets                 = "all"
  vrfs = {
    "VRF1" = {
      static_routes = {
        "2001:db8:3333:4444:5555:6666:102:304/128" = {
          control     = "bfd"
          description = "My Description"
          preference  = 10
          tag         = 100
          next_hops = {
            "unspecified;a:b::c:d/128;default" = {
              description           = "My Description"
              object                = 10
              preference            = 123
              tag                   = 10
              name                  = "nh-name"
              rewrite_encapsulation = "vlan-1"
            }
          }
        }
      }
      interfaces = {
        "eth1/10" = {
          auto_configuration         = "disabled"
          default_route              = "disabled"
          forward                    = "disabled"
          link_local_address_use_bia = "disabled"
          use_link_local_address     = "disabled"
          urpf                       = "disabled"
          link_local_address         = "2001:db8:3333:4444:5555:6666:7777:8888"
          addresses = {
            "2001:db8:3333:4444:5555:6666:7777:8888" = {
              type       = "primary"
              tag        = 1234
              control    = "anycast"
              preference = 10
              vpc_peer   = "2001:db8::1"
            }
          }
        }
      }
    }
  }
}
