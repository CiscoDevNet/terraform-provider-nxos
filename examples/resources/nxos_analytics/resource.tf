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
      monitors = {
        "MONITOR1" = {
          description = "My monitor"
        }
      }
      forward_instance_targets = {
        "1" = {
          default_policy               = "deny"
          collector_id                 = 1
          direction                    = "out"
          filter_type                  = "ipv6"
          instance_name                = "Ethernet1/1"
          switch_latency               = true
          system_exporter_id           = 201
          traffic_analytics_enabled    = true
          monitor_attachment_target_dn = "sys/analytics/inst-[analytics]/monitor-[MONITOR1]"
          profile_attachment_target_dn = "sys/analytics/inst-[analytics]/prof-[PROFILE1]"
          events_attachment_target_dn  = "sys/analytics/inst-[analytics]/events-[EVENTS1]"
          policy_attachment_target_dn  = "sys/analytics/inst-[analytics]/policy-[POLICY1]"
        }
      }
    }
  }
}
