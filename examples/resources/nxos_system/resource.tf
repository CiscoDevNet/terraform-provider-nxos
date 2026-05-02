resource "nxos_system" "example" {
  name                                               = "LEAF1"
  ethernet_mtu                                       = 9216
  ethernet_default_admin_state                       = "up"
  ethernet_admin_link_down_syslog_level              = 4
  ethernet_admin_link_up_syslog_level                = 4
  ethernet_admin_state                               = "enabled"
  ethernet_allow_unsupported_sfp                     = true
  ethernet_chassis_infrastructure_adaptor_vlan       = 100
  ethernet_chassis_infrastructure_epds_port_number   = 100
  ethernet_chassis_infrastructure_ipv6_address       = "2001:db8::1"
  ethernet_chassis_infrastructure_vlan               = 100
  ethernet_chassis_management_instance               = "mgmt0"
  ethernet_chassis_management_instance_fabric_number = "LeftFabric"
  ethernet_control                                   = "stateful-ha"
  ethernet_interface_syslog_info                     = "info-1"
  ethernet_log_event                                 = "linkStatusEnable"
  ethernet_default_layer                             = "Layer3"
  ethernet_system_interface_admin_state              = "up"
  ethernet_system_link_failure_laser_on              = false
  ethernet_system_storm_control_multi_threshold      = false
  ethernet_vlan_tag_native                           = false
  arp_admin_state                                    = "enabled"
  arp_instance_admin_state                           = "enabled"
  arp_allow_static_arp_outside_subnet                = "enabled"
  arp_unnumbered_svi_software_replication            = "enabled"
  arp_cache_limit                                    = 200000
  arp_cache_syslog_rate                              = 5
  arp_evpn_timeout                                   = 3000
  arp_interface_cache_limit                          = 1000
  arp_ip_adjacency_route_distance                    = 200
  arp_ip_arp_cos                                     = 4
  arp_off_list_timeout                               = 300
  arp_rarp_fabric_forwarding                         = "enabled"
  arp_rarp_fabric_forwarding_rate                    = 300
  arp_resolve_outside_subnet                         = "enabled"
  arp_suppression_timeout                            = 300
  arp_timeout                                        = 1800
  arp_vpc_domains = {
    "100" = {
    }
  }
  nd_admin_state                         = "enabled"
  nd_accept_solicit_neighbor_entry       = "accept"
  nd_instance_admin_state                = "enabled"
  nd_aging_interval                      = 1500
  nd_cache_limit                         = 200000
  nd_cache_syslog_rate                   = 5
  nd_ipv6_adjacency_route_distance       = 200
  nd_off_list_timeout                    = 300
  nd_probe_interval_for_solicit_neighbor = 10
  nd_solicit_neighbor_advertisement      = "enabled"
  nd_vrfs = {
    "default" = {
      interfaces = {
        "vlan100" = {
        }
      }
    }
  }
  nd_vpc_domains = {
    "100" = {
    }
  }
  clock_admin_state                = "enabled"
  clock_authentication_state       = "disabled"
  clock_format                     = "24hours"
  clock_format_debug               = false
  clock_format_syslog              = false
  clock_timezone_name              = "PST"
  clock_timezone_hours             = -8
  clock_timezone_minutes           = 0
  clock_summer_time_name           = "PDT"
  clock_summer_time_offset_minutes = 60
  clock_summer_time_start_week     = 2
  clock_summer_time_start_day      = "Sun"
  clock_summer_time_start_month    = "Mar"
  clock_summer_time_start_time     = "02:00"
  clock_summer_time_end_week       = 1
  clock_summer_time_end_day        = "Sun"
  clock_summer_time_end_month      = "Nov"
  clock_summer_time_end_time       = "02:00"
  dns_admin_state                  = "enabled"
  dns_profiles = {
    "default" = {
      description        = "My DNS Profile"
      owner_key          = "owner_key_1"
      owner_tag          = "owner_tag_1"
      domain_name        = "example.com"
      domain_description = "DNS domain description."
      domain_is_default  = false
    }
  }
  vdcs = {
    "1" = {
      name                                = "LEAF1"
      multicast_ipv4_route_memory_maximum = 50
      multicast_ipv4_route_memory_minimum = 50
      multicast_ipv6_route_memory_maximum = 5
      multicast_ipv6_route_memory_minimum = 5
      port_channel_maximum                = 500
      port_channel_minimum                = 10
      unicast_ipv4_route_memory_maximum   = 500
      unicast_ipv4_route_memory_minimum   = 500
      unicast_ipv6_route_memory_maximum   = 200
      unicast_ipv6_route_memory_minimum   = 200
      vlan_maximum                        = 4094
      vlan_minimum                        = 16
      vrf_maximum                         = 4096
      vrf_minimum                         = 4
    }
  }
  cli_aliases = {
    "myalias" = {
      command = "show ip route"
    }
  }
  smart_licensing_transport_mode     = "transportSmart"
  smart_licensing_transport_cslu_url = "https://cslu.example.com"
  boot_auto_copy                     = "disable"
  boot_dhcp                          = 500
  boot_exclude_configuration         = "enable"
  boot_order                         = "pxe"
  boot_poap                          = "enable"
  boot_image_verification            = "enable"
  cfs_admin_state                    = "enabled"
  udld_admin_state                   = "enabled"
  udld_aggressive                    = "enabled"
  udld_message_interval              = 20
  udld_interfaces = {
    "eth1/9" = {
      admin_state             = "port-enabled"
      aggressive              = "enabled"
      bidirectional_detection = "port-enabled"
    }
  }
  management_interfaces = {
    "mgmt0" = {
      admin_state      = "up"
      auto_negotiation = "on"
      description      = "My Description"
      duplex           = "auto"
      itu_channel      = 50
      media_type       = "auto"
      mtu              = 1500
      name             = "mgmt0"
      snmp_trap_state  = "enable"
      speed            = "auto"
    }
  }
  lldp_advertise_system_chassis_id = "enabled"
  lldp_hold_time                   = 180
  lldp_init_delay_time             = 5
  lldp_multi_peer                  = "disabled"
  lldp_optional_tlv_select         = "port-desc"
  lldp_port_channel                = "disabled"
  lldp_port_id_sub_type            = "short"
  lldp_system_description          = "Cisco NX-OS"
  lldp_transmit_frequency          = 60
  lldp_interfaces = {
    "eth1/1" = {
      admin_receive_state  = "disabled"
      admin_transmit_state = "disabled"
      port_dcbxp_version   = "CEE"
      port_description     = "My Port"
      system_description   = "Cisco NX-OS"
      tlv_management_ipv4  = "10.0.0.1"
      tlv_management_ipv6  = "2001:db8::1"
      tlv_vlan             = 100
    }
  }
  cdp_admin_state        = "enabled"
  cdp_device_id_type     = "mac"
  cdp_hold_interval      = 120
  cdp_pnp_startup_vlan   = 2
  cdp_transmit_frequency = 30
  cdp_version            = "v1"
  cdp_interfaces = {
    "eth1/1" = {
      admin_state      = "enabled"
      port_description = "My Port"
    }
  }
  copp_admin_state              = "enabled"
  copp_rate_limiter             = true
  console_exec_timeout          = 30
  vty_exec_timeout              = 30
  vty_session_limit             = 16
  icam_monitor_interval         = 4
  icam_number_of_intervals      = 336
  icam_scale_critical_threshold = 95
  icam_scale_info_threshold     = 70
  icam_scale_configuration      = true
  icam_scale_warning_threshold  = 85
  breakout_modules = {
    "1" = {
      front_panel_ports = {
        "49" = {
          breakout_map = "10g-4x"
        }
      }
    }
  }
  service_instances = {
    "hypershield" = {
      source_interface              = "lo100"
      controller_https_proxy_port   = 8080
      controller_https_proxy_server = "proxy.example.com"
      controller_ip1                = "10.0.0.1"
      controller_ip2                = "10.0.0.2"
      controller_ip3                = "10.0.0.3"
      firewall_policy_admin_state   = "in-service"
      vrfs = {
        "vrf1" = {
          affinity = 1
        }
      }
    }
  }
  ssh_admin_state                  = "enabled"
  ssh_ciphers                      = "no"
  ssh_description                  = "SSH access"
  ssh_enable_weak_ciphers          = "no"
  ssh_key_exchange_algorithms      = "no"
  ssh_key_types                    = "no"
  ssh_login_attempts               = 4
  ssh_login_grace_time             = 60
  ssh_message_authentication_codes = "no"
  ssh_port                         = 22
  ssh_keys = {
    "rsa" = {
      key_length = 2048
    }
  }
  erspan_origin_ip_is_global = true
  erspan_origin_ip_address   = "10.0.0.1"
  ttag_interfaces = {
    "eth1/10" = {
      ttag        = true
      ttag_inner  = false
      ttag_marker = false
      ttag_strip  = false
    }
  }
}
