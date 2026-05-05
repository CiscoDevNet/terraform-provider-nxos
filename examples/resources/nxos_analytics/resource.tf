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
      records = {
        "RECORD1" = {
          collect     = "count-bytes"
          description = "My record"
          match       = "src-ipv4"
        }
      }
      collectors = {
        "COLLECTOR1" = {
          description            = "My collector"
          dscp                   = 46
          destination_address    = "10.1.1.1"
          destination_port       = 5640
          event_destination_port = 5695
          inband_interface       = false
          source_address         = "10.1.1.2"
          source_interface       = "lo0"
          v9                     = false
          version                = "v9"
          vrf_name               = "default"
        }
      }
      monitors = {
        "MONITOR1" = {
          description      = "My monitor"
          record_target_dn = "sys/analytics/inst-analytics/recordp-RECORD1"
          collector_buckets = {
            "1" = {
              description = "My bucket"
              hash_high   = 4294967295
              hash_low    = 0
              collectors = {
                "sys/analytics/inst-analytics/collector-COLLECTOR1" = {}
              }
            }
          }
        }
      }
      traffic_analytics_description                  = "My traffic analytics"
      traffic_analytics_interface_mode               = true
      traffic_analytics_name                         = "TA1"
      traffic_analytics_service_database_size        = 1000
      traffic_analytics_troubleshoot_export_interval = 30
      traffic_analytics_udp_port_list                = "1234"
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
