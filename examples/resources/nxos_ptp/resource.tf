resource "nxos_ptp" "example" {
  domain_number               = 1
  notify_grandmaster_change   = "enabled"
  notify_parent_change        = "enabled"
  peer_delay_request_interval = 1
  priority1                   = 128
  priority2                   = 128
  using_system_clock          = true
  source_ip                   = "10.1.1.1"
  vrf_name                    = "default"
  domains = {
    "1" = {
    }
  }
  notify_high_correction_interval   = 10
  notify_high_correction            = "enabled"
  notify_high_correction_periodic   = "enabled"
  notify_port_state_change_category = "all"
  notify_port_state_change_interval = 10
  notify_port_state_change          = "enabled"
  notify_port_state_change_periodic = "enabled"
  interfaces = {
    "eth1/1" = {
      announce_interval_value          = 2
      announce_timeout_value           = 3
      delay_request_min_interval_value = 0
      negotiation_schema               = "schema1"
      ptp                              = true
      role                             = "master"
      sync_interval_value              = -1
      transmission                     = "multicast"
      transport                        = "ipv4"
      unicast_source                   = "10.2.2.1"
      unicast_vrf                      = "default"
      peers = {
        "10.1.1.2" = {
          negotiation_schema = "schema1"
        }
      }
    }
  }
}
