resource "nxos_dhcp" "example" {
  admin_state                                 = "enabled"
  ipv6_relay_information_option_vpn           = true
  ipv6_relay_option_type_cisco                = true
  relay_information_option                    = true
  relay_information_option_trust              = true
  relay_information_option_vpn                = true
  relay_information_trust_all                 = true
  relay_sub_option_circuit_id_customized      = true
  relay_sub_option_circuit_id_format_string   = "%p"
  relay_sub_option_type_cisco                 = true
  smart_relay_global                          = true
  snooping                                    = true
  snooping_information_option                 = true
  snooping_verify_mac_address                 = false
  dai_log_buffer_entries                      = 64
  dai_validate_destination                    = true
  dai_validate_ip                             = true
  dai_validate_source                         = true
  ipv6_relay_option79                         = true
  packet_strict_validation                    = true
  relay_dai                                   = true
  relay_information_option_server_id_override = 1
  relay_sub_option_format_non_tlv             = true
  relay_v4_over_v6                            = true
  relay_v6_iapd_route_add                     = true
  snoop_sub_option_circuit_id_format_string   = "%p"
  snooping_sub_option_format_non_tlv          = true
  v4_relay                                    = true
  v6_relay                                    = true
  v6_smart_relay_global                       = true
  relay_interfaces = [{
    interface_id        = "eth1/10"
    information_trusted = true
    smart_relay         = true
    subnet_broadcast    = true
    options             = "relay-info"
    subnet_selection    = "10.0.0.0"
    v6_smart_relay      = false
    addresses = [{
      vrf     = "VRF1"
      address = "1.1.1.1"
      counter = 1
    }]
  }]
}
