resource "nxos_eigrp" "example" {
  instances = {
    "EIGRP1" = {
      flush_routes = true
      isolate      = true
      vrfs = {
        "default" = {
          address_families = {
            "ipv4-ucast" = {
              admin_state                                                = "enabled"
              asn                                                        = 100
              router_id                                                  = "10.1.1.1"
              internal_distance                                          = 91
              external_distance                                          = 171
              maximum_paths                                              = 4
              active_interval_state                                      = "enabled"
              active_interval                                            = 5
              metric_style                                               = "wide"
              rib_scale                                                  = 100
              maximum_hops                                               = 50
              log_adjacency_changes                                      = false
              log_neighbor_warnings_interval                             = 20
              log_neighbor_warnings_state                                = "enabled"
              passive_interface_default                                  = true
              suppress_fib_pending                                       = true
              bfd                                                        = true
              authentication_mode                                        = "md5"
              authentication_key_chain                                   = "KEYCHAIN1"
              default_metric_bandwidth                                   = 10000
              default_metric_delay                                       = 200
              default_metric_reliability                                 = 200
              default_metric_load                                        = 2
              default_metric_mtu                                         = 1500
              graceful_restart                                           = false
              graceful_restart_route_hold_interval                       = 200
              graceful_restart_convergence_interval                      = 100
              graceful_restart_signal_interval                           = 30
              graceful_restart_await_redistribution_protocol_convergence = true
              metric_weights_type_of_service                             = 0
              metric_weights_k1                                          = 2
              metric_weights_k2                                          = 1
              metric_weights_k3                                          = 2
              metric_weights_k4                                          = 0
              metric_weights_k5                                          = 0
              metric_weights_k6                                          = 0
              default_information_originate_always                       = true
              default_information_originate_route_map                    = "RM_DEFAULT"
              redistributions = {
                "bgp;none;65001" = {
                  route_map = "RM_REDIST"
                }
              }
              redistribute_maximum_prefix_control   = "withdraw"
              redistribute_maximum_prefix           = 1000
              redistribute_maximum_prefix_threshold = 80
              redistribute_maximum_prefix_retries   = 2
              redistribute_maximum_prefix_duration  = 400
              table_map_route_map                   = "RM_TABLE"
              table_map_always                      = false
              stub_direct                           = true
              stub_static                           = true
              stub_summary                          = false
              stub_external                         = false
              stub_receive_only                     = false
              stub_leak_map                         = "RM_STUB"
            }
          }
          interfaces = {
            "eth1/10" = {
              address_families = {
                "ipv4-ucast" = {
                  admin_state              = "enabled"
                  hello_interval           = 10
                  hold_time                = 30
                  bandwidth                = 100000
                  bandwidth_percent        = 40
                  delay                    = 200
                  delay_unit               = "tens-of-micro"
                  mtu                      = 1400
                  bfd                      = "enabled"
                  next_hop_self            = false
                  split_horizon            = false
                  passive                  = "enabled"
                  authentication_mode      = "md5"
                  authentication_key_chain = "KEYCHAIN1"
                  distribute_lists = {
                    "in" = {
                      route_map = "RM_DIST"
                    }
                  }
                  offset_lists = {
                    "in" = {
                      route_map = "RM_OFFSET"
                      offset    = 10
                    }
                  }
                  summary_addresses = {
                    "10.10.0.0/16" = {
                      route_map = "RM_SUMMARY"
                      distance  = 10
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
