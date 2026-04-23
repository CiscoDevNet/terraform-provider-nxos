resource "nxos_port_channel_interface" "example" {
  port_channel_interfaces = {
    "po123" = {
      port_channel_mode                      = "active"
      minimum_links                          = 2
      maximum_links                          = 10
      suspend_individual                     = "disable"
      access_vlan                            = "vlan-1"
      admin_state                            = "up"
      auto_negotiation                       = "on"
      bandwidth                              = 0
      delay                                  = 1
      duplex                                 = "auto"
      layer                                  = "Layer2"
      link_logging                           = "enable"
      medium                                 = "broadcast"
      mode                                   = "access"
      mtu                                    = 1500
      native_vlan                            = "unknown"
      speed                                  = "auto"
      trunk_vlans                            = "1-4094"
      equalization_delay                     = 5
      graceful_convergence                   = "disable"
      hash_distribution                      = "adaptive"
      itu_channel                            = 50
      lacp_delay_mode                        = "enable"
      lacp_vpc_convergence                   = "enable"
      load_defer                             = "enable"
      mdix                                   = "auto"
      optics_loopback                        = "internal"
      pxe_transition_timeout                 = 5
      span_mode                              = "not-a-span-dest"
      squelch                                = "disable"
      transmission_mode                      = "not-a-trans-port"
      usage                                  = "discovery"
      user_configured_flags                  = "admin_layer,admin_state"
      storm_control_burst_packets_per_second = 600
      storm_control_burst_rate               = "75.000000"
      storm_control_rate                     = "50.000000"
      storm_control_rate_packets_per_second  = 500
      storm_control_packet_type              = "bcast"
      members = {
        "sys/intf/phys-[eth1/11]" = {
          force = true
        }
      }
    }
  }
}
