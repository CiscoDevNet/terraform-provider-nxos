resource "nxos_hardware_telemetry" "example" {
  admin_state                = "enabled"
  agent_address              = "172.24.141.69"
  counter_poll_interval      = 30
  control                    = "stateful-ha"
  extended_bgp               = true
  extended_switch            = true
  max_header_size            = 64
  packet_sampling_rate       = 4096
  receiver_address           = "10.92.198.113"
  receiver_max_datagram_size = 1400
  receiver_port              = 2055
  receiver_source_address    = "10.0.0.1"
  receiver_vrf_name          = "management"
}
