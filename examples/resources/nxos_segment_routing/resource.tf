resource "nxos_segment_routing" "example" {
  admin_state                    = "enabled"
  instance_admin_state           = "enabled"
  instance_control               = "stateful-ha"
  srgb_allocation_retry_interval = 120
  srgb_cleanup_interval          = 90
  global_blocks = {
    "mpls" = {
      label_blocks = {
        "srgb-1" = {
          minimum_label = 17000
          maximum_label = 25999
        }
      }
    }
  }
  domains = {
    "default" = {
      address_families = {
        "ipv4" = {
          name = "AF1"
          connected_prefix_sid_maps = {
            "mpls" = {
              prefix_sids = {
                "1.1.1.1/32" = {
                  sid_value      = 17001
                  sid_value_type = "absolute"
                }
              }
            }
          }
        }
      }
    }
  }
  traffic_engineering_admin_state = "enabled"
  traffic_engineering_encapsulations = {
    "mpls" = {
      source_address = "1.2.3.4"
    }
  }
  pcc_source_address = "1.1.1.1"
  pces = {
    "2.2.2.2" = {
      precedence = 10
    }
  }
  affinity_colors = {
    "RED" = {
      bit_position = 1
    }
  }
  liveness_detection_interval   = 1000
  liveness_detection_multiplier = 5
  traffic_engineering_interfaces = {
    "eth1/1" = {
      admin_state = "disabled"
      description = "My Description"
      name        = "IF1"
      affinity_colors = {
        "RED" = {}
      }
      metric = 10
    }
  }
  segment_lists = {
    "SL1" = {
      labels = {
        "10" = {
          label = 16001
        }
      }
      policy_endpoints = {
        "20" = {
          enable = true
        }
      }
    }
  }
  on_demand_colors = {
    "100" = {
      lockdown                       = true
      metric_type                    = "igp"
      shutdown                       = false
      liveness_detection_shutdown    = true
      liveness_detection_index_limit = 10
      candidate_paths = {
        "100" = {
          shutdown = true
          include_all_affinity_colors = {
            "RED" = {}
          }
          dataplane                = "mpls"
          protection_type          = "unprotected"
          dynamic_admin_state      = "enabled"
          dynamic_metric_type      = "te"
          dynamic_pcep_admin_state = "enabled"
        }
      }
    }
  }
  policies = {
    "P1" = {
      color                          = 200
      endpoint                       = "2.2.2.2"
      lockdown                       = true
      shutdown                       = true
      liveness_detection_shutdown    = true
      liveness_detection_index_limit = 10
      candidate_paths = {
        "100" = {
          shutdown = true
          explicit_segment_lists = {
            "SL1" = {}
          }
        }
      }
    }
  }
}
