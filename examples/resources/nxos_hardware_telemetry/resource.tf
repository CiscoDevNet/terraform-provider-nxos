resource "nxos_hardware_telemetry" "example" {
  sflow_admin_state                = "enabled"
  sflow_agent_address              = "172.24.141.69"
  sflow_counter_poll_interval      = 30
  sflow_control                    = "stateful-ha"
  sflow_extended_bgp               = true
  sflow_extended_switch            = true
  sflow_max_header_size            = 64
  sflow_packet_sampling_rate       = 4096
  sflow_receiver_address           = "10.92.198.113"
  sflow_receiver_max_datagram_size = 1400
  sflow_receiver_port              = 2055
  sflow_receiver_source_address    = "10.0.0.1"
  sflow_receiver_vrf_name          = "management"
}
