resource "nxos_logging" "example" {
  all   = "unspecified"
  level = "information"
  facilities = {
    "spanning-tree" = {
      level = "information"
    }
  }
  file_admin_state          = "disabled"
  file_description          = "syslog file"
  file_name                 = "messages"
  file_persistent_threshold = 90
  file_size                 = 4194304
  remote_destinations = {
    "10.0.0.1" = {
      admin_state                = "enabled"
      description                = "remote syslog server"
      name                       = "server1"
      port                       = 514
      transport                  = "udp"
      trustpoint_client_identity = "my-trustpoint"
      vrf_name                   = "management"
      severity                   = "warnings"
      forwarding_facility        = "local7"
    }
  }
  source_interface_admin_state = "enabled"
  source_interface_name        = "lo0"
  timestamp_format             = "milliseconds"
  monitor_admin_state          = "enabled"
  monitor_severity             = "warnings"
  console_admin_state          = "enabled"
  console_severity             = "warnings"
  origin_id_type               = "string"
  origin_id_value              = "SWITCH1"
}
