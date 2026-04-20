resource "nxos_bfd" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  instance_control     = "stateful-ha"
  echo_interface       = "lo10"
  hardware_offload     = "enable"
  slow_interval        = 3000
  startup_interval     = 10
  interfaces = {
    "vlan10" = {
      admin_state                 = "enabled"
      control                     = "opt-subif"
      destination                 = "10.1.1.1"
      echo_admin_state            = "disabled"
      source_ip                   = "10.1.1.2"
      start_timeout               = 10
      track_member_link           = "enable"
      vpc_watch                   = "enable"
      detect_multiplier           = 5
      echo_receive_interval       = 100
      min_receive_interval        = 100
      min_transmit_interval       = 100
      authentication_interop      = "enable"
      authentication_hex_key      = "AB"
      authentication_hex_key_size = 1
      authentication_key          = "secret_key"
      authentication_key_id       = 5
      authentication_type         = "sha1"
    }
  }
}
