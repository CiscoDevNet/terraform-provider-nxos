resource "nxos_bgp" "example" {
  admin_state                              = "enabled"
  instance_admin_state                     = "enabled"
  asn                                      = "65001"
  control                                  = "stateful-ha"
  disable_policy_batching                  = "disabled"
  disable_policy_batching_nexthop          = "disabled"
  disable_policy_batching_ipv4_prefix_list = "PFX_LIST1"
  disable_policy_batching_ipv6_prefix_list = "PFX_LIST_V6"
  enhanced_error_handling                  = false
  flush_routes                             = "disabled"
  isolate                                  = "disabled"
  nexthop_suppress_default_resolution      = "disabled"
  rd_dual                                  = "enabled"
  vrfs = {
    "default" = {
      router_id                          = "1.1.1.1"
      bestpath_first_always              = "disabled"
      bestpath_interval                  = 300
      bandwidth_reference                = 40000
      bandwidth_reference_unit           = "mbps"
      cluster_id                         = "1.2.3.4"
      hold_time                          = 180
      keepalive_interval                 = 60
      mode                               = "fabric"
      prefix_peer_timeout                = 90
      prefix_peer_wait_time              = 90
      reconnect_interval                 = 60
      router_id_auto                     = "disabled"
      route_control_enforce_first_as     = "disabled"
      route_control_fib_accelerate       = "enabled"
      route_control_log_neighbor_changes = "enabled"
      route_control_suppress_routes      = "disabled"
      graceful_restart_control           = "complete"
      graceful_restart_interval          = 240
      graceful_restart_stale_interval    = 1800
      address_families = {
        "ipv4-ucast" = {
          critical_nexthop_timeout               = "2500"
          non_critical_nexthop_timeout           = "8000"
          advertise_l2vpn_evpn                   = "disabled"
          advertise_physical_ip_for_type5_routes = "disabled"
          max_ecmp_paths                         = 2
          max_external_ecmp_paths                = 1
          max_external_internal_ecmp_paths       = 1
          max_local_ecmp_paths                   = 1
          max_mixed_ecmp_paths                   = 1
          default_information_originate          = "disabled"
          next_hop_route_map_name                = "ROUTEMAP1"
          prefix_priority                        = "none"
          retain_rt_all                          = "disabled"
          advertise_only_active_routes           = "disabled"
          table_map_route_map_name               = "ROUTE_MAP1"
          vni_ethernet_tag                       = "disabled"
          wait_igp_converged                     = "disabled"
          advertise_system_mac                   = "disabled"
          bestpath_origin_as_allow_invalid       = "disabled"
          bestpath_origin_as_use_validity        = "disabled"
          client_to_client_reflection            = "enabled"
          default_metric                         = "100"
          export_gateway_ip                      = "disabled"
          igp_metric                             = 600
          max_path_unequal_cost                  = "disabled"
          nexthop_load_balance_egress_multisite  = "disabled"
          origin_as_validate                     = "disabled"
          origin_as_validate_signal_ibgp         = "disabled"
          table_map_filter                       = "disabled"
          advertised_prefixes = {
            "192.168.1.0/24" = {
              route_map = "rt-map"
              evpn      = "enabled"
            }
          }
          redistributions = {
            "ospf;OSPF1" = {
              route_map        = "route_map_ospf_1"
              scope            = "inter"
              srv6_prefix_type = "unspecified"
            }
          }
          additional_paths_capability = "send"
          additional_paths_route_map  = "ADD_PATH_MAP"
          aggregate_addresses = {
            "10.0.0.0/8" = {
              advertise_map = "ADV_MAP"
              as_set        = "enabled"
              attribute_map = "ATTR_MAP"
              summary_only  = "enabled"
            }
          }
        }
      }
      peer_templates = {
        "SPINE-PEERS" = {
          remote_asn                     = "65002"
          description                    = "My Description"
          peer_type                      = "fabric-internal"
          source_interface               = "lo0"
          admin_state                    = "enabled"
          asn_type                       = "none"
          bfd_type                       = "none"
          bmp_server_1                   = "disabled"
          bmp_server_2                   = "disabled"
          capability_suppress_4_byte_asn = "disabled"
          connection_mode                = "passive"
          peer_control                   = "bfd"
          hold_time                      = 180
          keepalive_interval             = 60
          log_neighbor_changes           = "none"
          low_memory_exempt              = "disabled"
          private_as_control             = "none"
          session_template               = "SESS_TEMPLATE"
          peer_template_address_families = {
            "ipv4-ucast" = {
              control                     = "nh-self,rr-client"
              send_community_extended     = "enabled"
              send_community_standard     = "enabled"
              aigp                        = "disabled"
              allowed_self_as_count       = 0
              as_override                 = "disabled"
              default_originate           = "disabled"
              default_originate_route_map = "DEF_ORIG_MAP"
              dmz_link_bandwidth          = "disabled"
              link_bandwidth_cumulative   = "disabled"
              nexthop_thirdparty          = "enabled"
              rewrite_rt_asn              = "disabled"
              soft_reconfiguration_backup = "none"
              unsuppress_map              = "UNSUPP_MAP"
              weight                      = "100"
              max_prefix_action           = "restart"
              max_prefix_number           = 10000
              max_prefix_restart_time     = 1
              max_prefix_threshold        = 30
            }
          }
        }
      }
      peers = {
        "192.168.0.1" = {
          remote_asn                     = "65002"
          description                    = "My description"
          peer_template                  = "SPINE-PEERS"
          peer_type                      = "fabric-internal"
          source_interface               = "lo0"
          hold_time                      = 45
          keepalive_interval             = 15
          ebgp_multihop_ttl              = 5
          peer_control                   = "bfd,dis-conn-check"
          password                       = "secret_password"
          admin_state                    = "enabled"
          asn_type                       = "none"
          bfd_type                       = "none"
          bmp_server_1                   = "disabled"
          bmp_server_2                   = "disabled"
          capability_suppress_4_byte_asn = "disabled"
          connection_mode                = "passive"
          log_neighbor_changes           = "none"
          low_memory_exempt              = "disabled"
          private_as_control             = "none"
          session_template               = "SESS_TEMPLATE"
          local_asn_propagation          = "no-prepend"
          local_asn                      = "65001"
          peer_address_families = {
            "ipv4-ucast" = {
              control                     = "nh-self"
              send_community_extended     = "enabled"
              send_community_standard     = "enabled"
              aigp                        = "disabled"
              allowed_self_as_count       = 0
              as_override                 = "disabled"
              default_originate           = "disabled"
              default_originate_route_map = "DEF_ORIG_MAP"
              dmz_link_bandwidth          = "disabled"
              link_bandwidth_cumulative   = "disabled"
              nexthop_thirdparty          = "enabled"
              rewrite_rt_asn              = "disabled"
              soft_reconfiguration_backup = "none"
              unsuppress_map              = "UNSUPP_MAP"
              weight                      = "100"
              max_prefix_action           = "restart"
              max_prefix_number           = 10000
              max_prefix_restart_time     = 1
              max_prefix_threshold        = 30
              route_controls = {
                "in" = {
                  route_map_name = "ROUTE_MAP1"
                }
              }
              prefix_list_controls = {
                "in" = {
                  list = "PREFIX_LIST1"
                }
              }
            }
          }
        }
      }
      interface_peers = {
        "eth1/1" = {
          remote_asn                     = "65002"
          description                    = "My description"
          peer_template                  = "SPINE-PEERS"
          peer_type                      = "fabric-internal"
          admin_state                    = "enabled"
          asn_type                       = "none"
          bfd_type                       = "none"
          bmp_server_1                   = "disabled"
          bmp_server_2                   = "disabled"
          capability_suppress_4_byte_asn = "disabled"
          connection_mode                = "passive"
          peer_control                   = "bfd"
          dscp                           = "af11"
          dynamic_route_map              = "DYN_RT_MAP"
          hold_time                      = 45
          keepalive_interval             = 15
          log_neighbor_changes           = "none"
          low_memory_exempt              = "disabled"
          password                       = "secret_password"
          private_as_control             = "none"
          session_template               = "SESS_TEMPLATE"
          internal_vpn_client            = "disabled"
        }
      }
    }
  }
}
