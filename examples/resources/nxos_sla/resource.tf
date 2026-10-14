resource "nxos_sla" "example" {
  sender_logging_level = 5
  sender_logging_traps = true
  probes = {
    "1" = {
      icmp_echo_buckets_kept                     = 30
      icmp_echo_destination_ip                   = "10.1.1.1"
      icmp_echo_distributions_of_statistics_kept = 5
      icmp_echo_enhanced_buckets                 = 50
      icmp_echo_enhanced_interval                = 900
      icmp_echo_filter                           = "All"
      icmp_echo_frequency                        = 30
      icmp_echo_hours_of_statistics_kept         = 5
      icmp_echo_tos                              = 10
      icmp_echo_lives_kept                       = 1
      icmp_echo_owner                            = "OWNER1"
      icmp_echo_request_data_size                = 100
      icmp_echo_source_interface                 = "eth1/10"
      icmp_echo_statistics_distribution_interval = 10
      icmp_echo_tag                              = "TAG1"
      icmp_echo_threshold                        = 1000
      icmp_echo_timeout                          = 2000
      icmp_echo_verify_data                      = true
      icmp_echo_vrf                              = "management"
      reactions = {
        "rtt" = {
          lower_threshold   = 1000
          reaction_type     = "trapAndTrigger"
          threshold_count_x = 3
          threshold_count_y = 10
          threshold_type    = "xOfy"
          upper_threshold   = 2000
        }
      }
      reaction_triggers = {
        "2" = {}
      }
      schedule_age_out     = 86400
      schedule_life        = 7200
      schedule_recurring   = true
      schedule_start_probe = "At"
      schedule_start_time  = "10:00:00"
    }
  }
  group_schedules = {
    "1" = {
      age_out         = 3600
      frequency       = 30
      life            = 7200
      probe_list      = "3"
      schedule_period = 10
      start_probe     = "After"
      start_time      = "01:00:00"
    }
  }
  responder_admin_state   = true
  responder_logging_level = 5
  permanent_responders = {
    "udp;5000;10.1.1.1" = {}
  }
  twamp_responder_admin_state   = true
  twamp_responder_timeout       = 600
  twamp_server_admin_state      = true
  twamp_server_inactivity_timer = 600
  twamp_server_logging_level    = 5
  twamp_server_port             = 5862
}
