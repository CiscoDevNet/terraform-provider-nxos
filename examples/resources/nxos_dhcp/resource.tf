resource "nxos_dhcp" "example" {
  admin_state                                                          = "enabled"
  instance_ipv6_relay_information_option_vpn_enabled                   = true
  instance_ipv6_relay_option_type_cisco_enabled                        = true
  instance_relay_information_option_enabled                            = true
  instance_relay_information_option_trust_enabled                      = true
  instance_relay_information_option_vpn_enabled                        = true
  instance_relay_information_trust_all_enabled                         = true
  instance_relay_sub_option_circuit_id_customized_enabled              = true
  instance_relay_sub_option_circuit_id_format_string                   = "%p"
  instance_relay_sub_option_type_cisco_enabled                         = true
  instance_smart_relay_global_enabled                                  = true
  instance_snooping_enabled                                            = true
  instance_snooping_information_option_enabled                         = true
  instance_snooping_verify_mac_address_enabled                         = false
  instance_dai_log_buffer_entries                                      = 64
  instance_dai_validate_destination                                    = true
  instance_dai_validate_ip                                             = true
  instance_dai_validate_source                                         = true
  instance_ipv6_relay_option79_enabled                                 = true
  instance_packet_strict_validation                                    = true
  instance_relay_dai_enabled                                           = true
  instance_relay_information_option_server_id_override_disable_enabled = 1
  instance_relay_sub_option_format_non_tlv_enabled                     = true
  instance_relay_v4_over_v6_enabled                                    = true
  instance_relay_v6_iapd_route_add_enabled                             = true
  instance_snoop_sub_option_circuit_id_format_string                   = "%p"
  instance_snooping_sub_option_format_non_tlv_enabled                  = true
  instance_v4_relay_enabled                                            = true
  instance_v6_relay_enabled                                            = true
  instance_v6_smart_relay_global_enabled                               = true
  relay_interfaces = [{
    interface_id                = "eth1/10"
    information_trusted_enabled = true
    smart_relay_enabled         = true
    subnet_broadcast_enabled    = true
    options                     = "relay-info"
    subnet_selection            = "10.0.0.0"
    v6_smart_relay_enabled      = false
    addresses = [{
      vrf     = "VRF1"
      address = "1.1.1.1"
      counter = 1
    }]
  }]
}
