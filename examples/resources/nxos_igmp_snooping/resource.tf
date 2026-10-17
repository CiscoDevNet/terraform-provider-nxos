resource "nxos_igmp_snooping" "example" {
  instance_admin_state                       = "enabled"
  instance_control                           = "stateful-ha"
  domain_control                             = "opt-flood"
  domain_name                                = "default"
  global_vlan_vxlan                          = true
  global_vlan_disable_nve_static_router_port = false
  global_vlan_vxlan_umc_drop_vlan            = "100-200"
  vlans = {
    "vlan3901" = {
      report_flood_all_interfaces   = true
      explicit_tracking             = false
      igmp_snooping                 = false
      igmp_source_ip_proxy          = false
      link_local_groups_suppression = false
      lookup_mac                    = true
      query_flood                   = false
      report_suppression            = false
      v3_report_suppression         = true
      querier_address               = "10.1.1.1"
      querier_description           = "My querier description"
      last_member_query_interval    = 2
      querier_name                  = "querier"
      query_interval                = 200
      robustness_variable           = 3
      query_max_response_time       = 20
      startup_query_count           = 3
      startup_query_interval        = 40
      querier_timeout               = 300
      version                       = "v3"
      fast_leave                    = true
      group_timeout                 = 300
      max_general_query_miss        = 4
      proxy_leave_use_group_address = true
      proxy_max_response_time       = 10
      access_group_description      = "My access group description"
      access_group_direction        = "in"
      access_group_name             = "access-group"
      access_group_route_map        = "RM1"
      mrouter_interfaces = {
        "eth1/10" = {}
      }
      static_groups = {
        "239.1.1.1" = {
          name = "group1"
          interfaces = {
            "eth1/11" = {}
          }
          sources = {
            "10.1.1.5" = {
              interfaces = {
                "eth1/11" = {}
              }
            }
          }
        }
      }
      report_flood_interfaces = {
        "eth1/12" = {
          description  = "My description"
          name         = "report-flood"
          report_flood = true
        }
      }
    }
  }
}
