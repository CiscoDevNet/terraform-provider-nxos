resource "nxos_platform" "example" {
  nve_interfaces = {
    "1" = {
      ipmc_index_size = 4000
      infra_vlans = {
        "100" = {
          force = "Enable"
        }
      }
    }
  }
  extended_dme_load_interval              = 300
  port_channel_load_balance_concatenation = "enable"
  port_channel_load_balance_fields        = "srcDstIp"
  port_channel_load_balance_rotate        = 3
  port_channel_load_balance_symmetry      = "enable"
  dlb_allow_all_interfaces                = true
  dlb_decay_factor                        = 4
  dlb_description                         = "My DLB"
  dlb_dre_threshold_level_1               = 40
  dlb_dre_threshold_level_2               = 35
  dlb_dre_threshold_level_3               = 30
  dlb_dre_threshold_level_4               = 25
  dlb_dre_threshold_level_5               = 20
  dlb_dre_threshold_level_6               = 15
  dlb_dre_threshold_level_7               = 10
  dlb_flowlet_aging                       = 1000
  dlb_interfaces                          = "eth1/10,eth1/11"
  dlb_load_awareness                      = false
  dlb_mac_address                         = "00:AA:BB:CC:DD:EE"
  dlb_mode                                = "PER_PACKET"
  dlb_sampling_interval                   = 64000
  dlb_static_pinning                      = true
  dlb_static_pinnings = {
    "eth1/10" = {
      destination_interface = "eth1/11"
    }
  }
  pfc_watchdog_auto_restore_multiplier  = 20
  pfc_watchdog_mode                     = "on"
  pfc_watchdog_fixed_restore_multiplier = 50
  pfc_watchdog_interval                 = 200
  pfc_watchdog_shutdown_multiplier      = 5
  qos_burst_detect_fall_threshold_unit  = "percentage"
  qos_burst_detect_fall_threshold       = 40
  qos_burst_detect_rise_threshold_unit  = "percentage"
  qos_burst_detect_rise_threshold       = 60
  qos_classify_ns_only                  = "Enabled"
  qos_cpu_size                          = 300
  qos_drop_pool_size                    = 10
  qos_headroom_size                     = 10
  qos_include_ipg                       = "Enabled"
  qos_lcpu_size                         = 300
  qos_nodrop_dynamic_queue_limit        = 10
  qos_nodrop_resume_offset              = 10
  qos_nodrop_pool_size                  = 10
  qos_pfc_static                        = "Enabled"
  qos_pg_headroom_reserve               = 50
  qos_pg_min_value                      = 5
  qos_pg_no_min                         = false
  qos_pg_share                          = "Disabled"
  tcam_templates = {
    "TPL1" = {
      reference_template = "l3"
      udfs = {
        "racl;UDF1" = {}
      }
    }
  }
  udf_description = "My UDFs"
  udfs = {
    "UDF1" = {
      mpls_stack_1    = true
      mpls_stack_2    = true
      netflow_rtp     = true
      offset_base     = "outer-header-l3"
      offset_length   = 2
      offset_position = 10
      qualify_racl    = true
    }
  }
}
