resource "nxos_pim" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  control              = "stateful-ha"
  evpn_border_leaf     = false
  extra_net            = false
  join_prune_delay     = 200
  register_stop        = false
  isolate              = true
  log_level            = 3
  syslog_level         = 3
  vrfs = {
    "default" = {
      admin_state          = "enabled"
      bfd                  = true
      auto_enable          = false
      control              = "flush-on-restart"
      flush_routes         = false
      join_prune_delay     = 200
      log_neighbor_changes = false
      mtu                  = 1600
      register_rate_limit  = 100
      rfc_strict           = false
      spt_switch_graceful  = false
      interfaces = {
        "eth1/10" = {
          admin_state               = "enabled"
          bfd                       = "enabled"
          dr_priority               = 10
          passive                   = false
          sparse_mode               = true
          border                    = false
          border_router             = false
          control                   = "border"
          description               = "MyDescription"
          dr_delay                  = 5
          join_prune_route_map      = "JP_POLICY"
          name                      = "pim-if"
          neighbor_route_map        = "NEIGH_POLICY"
          pfm_sd_boundary           = 0
          rfc_strict                = false
          hello_authentication_key  = "cisco123"
          hello_authentication_type = "ah-md5"
          hello_description         = "My hello policy"
          hello_interval            = 5000
          hello_name                = "hello"
        }
      }
      ssm_policy_name              = "SSM"
      ssm_policy_description       = "SSM_Policy"
      ssm_range_group_list_1       = "232.0.0.0/8"
      ssm_range_group_list_2       = "233.0.0.0/8"
      ssm_range_none               = false
      static_rp_policy_name        = "RP"
      static_rp_policy_description = "Static_RP_Policy"
      static_rps = {
        "1.2.3.4" = {
          group_lists = {
            "225.0.0.0/24" = {
              bidir    = true
              override = true
            }
          }
        }
      }
      anycast_rp_local_interface  = "eth1/10"
      anycast_rp_source_interface = "eth1/10"
      anycast_rp_description      = "Anycast_RP"
      anycast_rp_name             = "anycast-rp"
      anycast_rp_peers = {
        "10.1.1.1/32;20.1.1.1/32" = {}
      }
      asm_policy_description                = "My ASM policy"
      pre_build_spt_force                   = false
      asm_policy_name                       = "asm"
      pre_build_spt                         = true
      shared_tree_route_map                 = "SHARED_MAP"
      shared_tree_ssm_none                  = false
      shared_tree_use_spt_threshold_command = true
      sg_expiry_timer                       = 300
      sg_expiry_timer_sparse                = false
      sg_expiry_timer_ssm_none              = false
      register_policy_description           = "My register policy"
      register_policy_max_rate              = 1000
      register_policy_name                  = "register"
      register_policy_route_map             = "REG_MAP"
      register_source_interface             = "eth1/10"
      bsr_control                           = "forward,listen"
      bsr_description                       = "My BSR policy"
      bsr_name                              = "bsr"
      bsr_candidate_description             = "My BSR candidate"
      bsr_candidate_hash_length             = 30
      bsr_candidate_interval                = 30
      bsr_candidate_name                    = "bsr-candidate"
      bsr_candidate_priority                = 10
      bsr_candidate_source_interface        = "eth1/10"
      bsr_rp_candidate_description          = "My RP candidate"
      bsr_rp_candidate_interval             = 30
      bsr_rp_candidate_name                 = "rp-candidate"
      bsr_rp_candidate_priority             = 10
      bsr_rp_candidate_source_interface     = "eth1/10"
      bsr_rp_candidate_group_ranges = {
        "no" = {
          group_list_1 = "239.0.0.0/8"
          override     = false
          ssm_none     = false
        }
      }
      bsr_policy_description                              = "My policy"
      bsr_policy_name                                     = "policy"
      bsr_policy_route_map                                = "BSR_POLICY"
      bsr_rp_candidate_policy_description                 = "My policy"
      bsr_rp_candidate_policy_name                        = "policy"
      bsr_rp_candidate_policy_route_map                   = "BSR_RP_POLICY"
      auto_rp_control                                     = "forward,listen"
      auto_rp_description                                 = "My Auto-RP policy"
      auto_rp_name                                        = "auto-rp"
      auto_rp_mapping_agent_description                   = "My mapping agent"
      auto_rp_mapping_agent_name                          = "mapping-agent"
      auto_rp_mapping_agent_scope                         = 16
      auto_rp_mapping_agent_source_interface              = "eth1/10"
      auto_rp_mapping_agent_use_send_rp_discovery_command = false
      auto_rp_rp_candidate_description                    = "My RP candidate"
      auto_rp_rp_candidate_interval                       = 30
      auto_rp_rp_candidate_name                           = "rp-candidate"
      auto_rp_rp_candidate_scope                          = 16
      auto_rp_rp_candidate_source_interface               = "eth1/10"
      auto_rp_rp_candidate_use_send_rp_announce_command   = false
      auto_rp_rp_candidate_group_ranges = {
        "no" = {
          group_list_1 = "239.0.0.0/8"
          override     = false
          ssm_none     = false
        }
      }
      auto_rp_mapping_agent_policy_description = "My policy"
      auto_rp_mapping_agent_policy_name        = "policy"
      auto_rp_mapping_agent_policy_route_map   = "MA_POLICY"
      auto_rp_rp_candidate_policy_description  = "My policy"
      auto_rp_rp_candidate_policy_name         = "policy"
      auto_rp_rp_candidate_policy_route_map    = "AUTO_RP_POLICY"
    }
  }
}
