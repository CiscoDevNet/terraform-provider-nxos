resource "nxos_isis" "example" {
  admin_state = "enabled"
  instances = {
    "ISIS1" = {
      admin_state  = "enabled"
      control      = "stateful-ha"
      flush_routes = true
      isolate      = true
      vrfs = {
        "default" = {
          admin_state              = "enabled"
          authentication_check_l1  = false
          authentication_check_l2  = false
          authentication_key_l1    = ""
          authentication_key_l2    = ""
          authentication_type_l1   = "unknown"
          authentication_type_l2   = "unknown"
          bandwidth_reference      = 400000
          bandwidth_reference_unit = "mbps"
          is_type                  = "l2"
          metric_type              = "wide"
          mtu                      = 2000
          net                      = "49.0001.0000.0000.3333.00"
          passive_default          = "l12"
          control                  = "log-adj-changes"
          lsp_lifetime             = 1000
          queue_limit              = 3000
          address_families = {
            "v4" = {
              segment_routing_mpls                    = true
              enable_bfd                              = false
              prefix_advertise_passive_l1             = true
              prefix_advertise_passive_l2             = true
              control                                 = "adj-check"
              default_information_originate           = "on"
              default_information_originate_route_map = "rm1"
              distance                                = 100
              max_ecmp                                = 4
              multi_topology                          = "st"
              router_id_interface                     = "lo0"
              table_map                               = "rm1"
              table_map_filter                        = "enabled"
              redistributions = {
                "direct;none;none" = {
                  route_map        = "RM_REDIST"
                  scope            = "inter"
                  srv6_prefix_type = "unspecified"
                }
              }
              summary_addresses = {
                "10.0.0.0/8" = {
                  level = "l2"
                }
              }
              distributions = {
                "l2" = {
                  destination_level = "l1"
                  route_map         = "RM_DISTRIBUTE"
                }
              }
            }
          }
          levels = {
            "l2" = {
              control                         = "fast-converge"
              spf_initial_interval            = 60
              spf_max_interval                = 5000
              spf_second_interval             = 60
              lsp_generation_initial_interval = 70
              lsp_generation_max_interval     = 6000
              lsp_generation_second_interval  = 70
            }
          }
          graceful_restart_admin_state  = "disabled"
          graceful_restart_t3_interval  = 120
          overload_startup_time         = 60
          overload_admin_state          = "always-on"
          overload_bgp_as_number        = 100
          overload_bgp_as_number_string = "100"
          overload_suppress             = "interlevel"
        }
      }
    }
  }
  interfaces = {
    "eth1/10" = {
      authentication_check         = false
      authentication_check_l1      = false
      authentication_check_l2      = false
      authentication_key           = ""
      authentication_key_l1        = ""
      authentication_key_l2        = ""
      authentication_type          = "unknown"
      authentication_type_l1       = "unknown"
      authentication_type_l2       = "unknown"
      circuit_type                 = "l2"
      vrf                          = "default"
      hello_interval               = 20
      hello_interval_l1            = 20
      hello_interval_l2            = 20
      hello_multiplier             = 4
      hello_multiplier_l1          = 4
      hello_multiplier_l2          = 4
      hello_padding                = "never"
      instance_name                = "ISIS1"
      metric_l1                    = 1000
      metric_l2                    = 1000
      mtu_check                    = true
      mtu_check_l1                 = true
      mtu_check_l2                 = true
      network_type_p2p             = "on"
      passive                      = "l1"
      priority_l1                  = 80
      priority_l2                  = 80
      enable_ipv4                  = true
      admin_state                  = "enabled"
      csnp_interval_l1             = 30
      csnp_interval_l2             = 30
      control                      = "advert-tep"
      description                  = "ISIS interface"
      lsp_refresh_interval         = 100
      mesh_group_id                = 10
      ipv6_metric_l1               = 1000
      ipv6_metric_l2               = 1000
      n_flag_clear                 = true
      retransmit_interval          = 10
      retransmit_throttle_interval = 100
      suppressed_state             = true
      ipv4_bfd                     = "enabled"
      ipv6_bfd                     = "enabled"
      ipv6                         = true
    }
  }
}
