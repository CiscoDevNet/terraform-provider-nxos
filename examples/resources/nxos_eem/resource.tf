resource "nxos_eem" "example" {
  environment_variables = {
    "TF_INTERFACE" = {
      value = "eth1/10"
    }
  }
  applets = {
    "TF_APPLET" = {
      description               = "Terraform applet"
      correlation_tag_1         = "TF_CLI"
      correlation_operator_1    = "and"
      correlation_tag_2         = "TF_SYSLOG"
      correlation_operator_2    = "and"
      correlation_tag_3         = "TF_COUNTER"
      correlation_operator_3    = "and"
      correlation_tag_4         = "TF_SNMP"
      correlation_occurrences   = 2
      correlation_time_interval = 60
      cli_events = {
        "TF_CLI" = {
          match = "show terraform"
        }
      }
      counter_events = {
        "TF_COUNTER" = {
          name           = "TF_COUNTER"
          entry_operator = "gt"
          entry_value    = 5
          exit_operator  = "lt"
          exit_value     = 2
        }
      }
      snmp_events = {
        "TF_SNMP" = {
          oid                   = "1.3.6.1.2.1.1.3.0"
          get_type              = "exact"
          entry_operator        = "lt"
          entry_value           = 1
          exit_operator         = "gt"
          exit_value            = 2
          poll_interval_minimum = 30
          poll_interval_maximum = 60
        }
      }
      syslog_events = {
        "TF_SYSLOG" = {
          pattern  = "terraform-test-pattern"
          priority = "alerts"
        }
      }
      actions = {
        "1.0" = {
          cli_command = "show clock"
          cli_local   = true
        }
        "2.0" = {
          counter_name      = "TF_COUNTER"
          counter_operation = "inc"
          counter_value     = "1"
        }
        "3.0" = {
          event_default = true
        }
        "4.0" = {
          exception_log_module      = "1"
          exception_log_syserr      = "0x1"
          exception_log_device_id   = "1"
          exception_log_error_type  = "1"
          exception_log_error_code  = "1"
          exception_log_phy_layer   = "1"
          exception_log_ports       = "1"
          exception_log_hard_error  = "soft_error"
          exception_log_description = "Terraform"
        }
        "5.0" = {
          publish_event_subsystem_id = 798
          publish_event_type         = 1
          publish_event_argument_1   = "a"
          publish_event_argument_2   = "b"
          publish_event_argument_3   = "c"
          publish_event_argument_4   = "d"
        }
        "6.0" = {
          snmp_trap                = true
          snmp_trap_integer_data_1 = "1"
          snmp_trap_integer_data_2 = "2"
          snmp_trap_string_data    = "Terraform"
        }
        "7.0" = {
          syslog_priority = "alerts"
          syslog_message  = "Terraform"
        }
      }
    }
    "TF_APPLICATION" = {
      application_event_subsystem_id = 798
      application_event_type         = 1
    }
    "TF_CDP" = {
      cdp_event_interface = "eth1/10"
      cdp_event_operation = "add"
    }
    "TF_CLI" = {
      cli_event_match         = "show terraform"
      cli_event_occurrences   = 2
      cli_event_time_interval = 10
    }
    "TF_COUNTER" = {
      counter_event_name           = "TF_COUNTER"
      counter_event_entry_operator = "gt"
      counter_event_entry_value    = 5
      counter_event_exit_operator  = "lt"
      counter_event_exit_value     = 2
    }
    "TF_ETHPM" = {
      ethpm_event_storm_control = true
    }
    "TF_FAN_ABSENT" = {
      fan_absent_event_fan  = "1"
      fan_absent_event_time = 20
    }
    "TF_FAN_BAD" = {
      fan_bad_event_fan  = "1"
      fan_bad_event_time = 20
    }
    "TF_GOLD" = {
      gold_event_module               = 1
      gold_event_test_name            = "TestUSBPorts"
      gold_event_testing_type         = "scheduled"
      gold_event_severity             = "major"
      gold_event_consecutive_failures = 2
    }
    "TF_LINK_FLAP" = {
      internal_link_flap_event_count = 2
      internal_link_flap_event_time  = 60
    }
    "TF_MEMORY" = {
      memory_event_level = "critical"
    }
    "TF_MODULE" = {
      module_event_module = "1"
      module_event_status = "online"
    }
    "TF_MODULE_FAILURE" = {
      module_failure_event_module     = "1"
      module_failure_event_error_type = "lc-failed"
      module_failure_event_count      = 2
      module_failure_event_time       = 60
    }
    "TF_NONE" = {
      none_event = true
    }
    "TF_OIR" = {
      oir_event_device_type   = "fan"
      oir_event_type          = "remove"
      oir_event_device_number = 1
    }
    "TF_POLICY_DEFAULT" = {
      override                           = "__pfm_fanbad_any_singlefan"
      policy_default_event_occurrences   = 2
      policy_default_event_time_interval = 60
    }
    "TF_POWER_OVER_BUDGET" = {
      power_over_budget_event = true
    }
    "TF_SNMP" = {
      snmp_event_oid                   = "1.3.6.1.2.1.1.3.0"
      snmp_event_get_type              = "exact"
      snmp_event_entry_operator        = "lt"
      snmp_event_entry_value           = 1
      snmp_event_exit_operator         = "gt"
      snmp_event_exit_value            = 2
      snmp_event_poll_interval_minimum = 30
      snmp_event_poll_interval_maximum = 60
    }
    "TF_SYSLOG" = {
      syslog_event_pattern       = "terraform-test-pattern"
      syslog_event_priority      = "alerts"
      syslog_event_occurrences   = 2
      syslog_event_time_interval = 60
    }
    "TF_SYSMGR_MEMORY" = {
      override                        = "__sysmgr_policy_mem_alert"
      sysmgr_memory_alert_event_minor = 70
      sysmgr_memory_alert_event_major = 80
      sysmgr_memory_alert_event_clear = 60
    }
    "TF_SYSMGR_SWITCHOVER" = {
      override                      = "__sysmgr_swover_count_alert"
      sysmgr_switchover_event_count = 3
      sysmgr_switchover_event_time  = 120
    }
    "TF_TEMPERATURE" = {
      temperature_event_module    = "1"
      temperature_event_sensor    = "1"
      temperature_event_threshold = "major"
    }
    "TF_TIMER" = {
      timer_event_name      = "TF_TIMER"
      timer_event_cron_time = "0 * * * *"
    }
    "TF_TRACK" = {
      track_event_id    = 10
      track_event_state = "down"
    }
    "TF_UFDM" = {
      ufdm_event_fib = "inconsistent"
    }
  }
}
