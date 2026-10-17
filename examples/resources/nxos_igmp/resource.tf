resource "nxos_igmp" "example" {
  instance_admin_state    = "enabled"
  any_query_destination   = true
  bootup_delay            = 10
  control                 = "stateful-ha"
  flush_routes            = true
  heavy_template          = true
  instance_control        = "rtr-alert"
  enforce_router_alert    = true
  spoof_check             = false
  strict_querier_election = true
  syslog_threshold        = 80
  vrfs = {
    "default" = {
      ssm_translations = {
        "232.1.1.0/24;10.3.3.3/32" = {}
      }
      interfaces = {
        "eth1/10" = {
          admin_state                     = "enabled"
          allow_v3_asm                    = false
          description                     = "My description"
          group_timeout                   = 300
          control                         = "fast-leave"
          immediate_leave                 = true
          name                            = "igmp-if"
          report_link_local_groups        = true
          suppress_v3_gsq                 = true
          version                         = "v3"
          querier_address                 = "10.1.1.1"
          querier_description             = "My querier description"
          last_member_query_count         = 3
          last_member_query_response_time = 5
          querier_name                    = "querier"
          querier_timeout                 = 300
          query_interval                  = 100
          robustness_variable             = 3
          query_max_response_time         = 20
          startup_query_count             = 4
          startup_query_interval          = 40
          querier_version                 = "unspecified"
          report_policy_route_map         = "REPORT_POLICY"
          report_policy_access_group      = true
          state_limit_description         = "My state limit description"
          state_limit_maximum             = 100
          state_limit_name                = "state-limit"
          state_limit_reserved            = 50
          state_limit_reserved_route_map  = "RESERVED_RM"
          join_types = {
            "1" = {
              versions = {
                "3" = {
                  groups = {
                    "239.1.1.2" = {
                      use_static_oif = false
                      sources = {
                        "10.1.1.1" = {
                          use_static_oif = false
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
  }
}
