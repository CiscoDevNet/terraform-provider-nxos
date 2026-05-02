resource "nxos_analytics" "example" {
  instances = {
    "analytics" = {
      admin_state              = "enabled"
      collect_tunnel_header    = true
      control                  = "stateful-ha"
      enable_analytics_submode = true
      geneve_enable            = true
      timeout                  = 30
      profiles = {
        "PROFILE1" = {
          burst_interval_shift                 = 10
          collect_interval                     = 200
          description                          = "My profile"
          ip_packet_id_shift                   = 10
          mtu                                  = 1500
          sequence_number_guess_threshold_high = 100
          sequence_number_guess_threshold_low  = 50
          source_port                          = 10000
        }
      }
      events = {
        "EVENTS1" = {
          acl_drops              = true
          black_hole             = true
          buffer_drops           = true
          description            = "My events"
          event_export_max       = 10
          forward_drops          = true
          group_drop_events      = true
          group_latency_events   = true
          group_packet_events    = true
          ip_dont_fragment       = true
          latency_threshold      = 100
          latency_threshold_unit = "milli-sec"
          receive_window_zero    = true
          tos                    = 10
          tos_enable             = true
          ttl_match_enable       = true
          ttl_match_value        = 10
        }
      }
      policies = {
        "POLICY1" = {
          description = "My policy"
        }
      }
      traffic_analytics_description                  = "My traffic analytics"
      traffic_analytics_interface_mode               = true
      traffic_analytics_name                         = "TA1"
      traffic_analytics_service_database_size        = 1000
      traffic_analytics_troubleshoot_export_interval = 30
      traffic_analytics_udp_port_list                = "1234"
    }
  }
}
