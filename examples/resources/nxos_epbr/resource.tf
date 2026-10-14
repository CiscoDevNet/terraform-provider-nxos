resource "nxos_epbr" "example" {
  services = {
    "SVC1" = {
      admin_state                    = "shut"
      connection_mode                = "full-duplex"
      hold_down_threshold_count      = 3
      hold_down_threshold_time       = 100
      reverse_security_group         = 200
      security_group                 = 100
      service_interface              = "eth1/10"
      threshold_up                   = 50
      type                           = "l3"
      vrf                            = "default"
      probe_control                  = "enable"
      probe_frequency                = 5
      probe_port                     = 80
      probe_protocol                 = "TCP"
      probe_retry_down_count         = 2
      probe_retry_up_count           = 2
      probe_reverse_source_interface = "lo123"
      probe_source_interface         = "lo123"
      probe_timeout                  = 3
      endpoints = {
        "10.1.1.10" = {
          admin_state               = "shut"
          hold_down_threshold_count = 3
          hold_down_threshold_time  = 100
          mode                      = "hot-standby"
          weight                    = 3
          probe_frequency           = 60
          probe_http_get            = "index.html"
          probe_http_version        = "1.1"
          probe_protocol            = "HTTP"
          probe_retry_down_count    = 2
          probe_retry_up_count      = 2
          probe_source_interface    = "lo123"
          probe_timeout             = 30
          reverse_ip                = "10.2.2.10"
          reverse_probe_protocol    = "noProtocol"
          reverse_probe_track_id    = 10
        }
      }
    }
  }
  failover_groups = {
    "FG1" = {
      service_chains = {
        "CHAIN2" = {
          preference = 10
        }
      }
    }
  }
  service_chains = {
    "CHAIN1" = {
      failover_group      = "FG1"
      load_balance_method = "src"
      mode                = "multisite"
      sequences = {
        "10" = {
          service             = "SVC2"
          action              = "route"
          load_balance_method = "dst"
          reverse_action      = "route"
        }
      }
    }
  }
  policies = {
    "POL1" = {
      default_traffic_action = "redirect"
      destination_vrf        = "default"
      source_vrf             = "default"
      statistics             = true
      matches = {
        "EPBR-ACL1" = {
          access_list_type    = "ipv4"
          buckets             = 8
          destination_vrf     = "default"
          load_balance_method = "src"
          mask_position       = 2
          source_vrf          = "default"
          traffic_action      = "redirect"
          sequences = {
            "10" = {
              service     = "SVC3"
              fail_action = "drop"
            }
          }
        }
      }
    }
  }
  infra_vlans = {
    "100" = {}
  }
  interfaces = {
    "eth1/11;ipv4;no-reverse" = {
      policy = "POL1"
    }
  }
}
