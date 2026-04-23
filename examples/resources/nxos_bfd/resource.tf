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
      admin_state            = "enabled"
      echo_admin_state       = "disabled"
      start_timeout          = 10
      detect_multiplier      = 5
      echo_receive_interval  = 100
      min_receive_interval   = 100
      min_transmit_interval  = 100
      authentication_interop = "enable"
      authentication_key     = "secret_key"
      authentication_key_id  = 5
      authentication_type    = "sha1"
    }
  }
}
