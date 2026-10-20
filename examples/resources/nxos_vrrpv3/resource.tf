resource "nxos_vrrpv3" "example" {
  admin_state = "enabled"
  interfaces = {
    "vlan10" = {
      delay_minimum = 5
      reload_delay  = 10
      groups = {
        "20;ipv4" = {
          admin_state               = "enabled"
          advertisement_interval    = 2000
          decrement_track_priority  = 20
          primary_address           = "10.1.1.10"
          match_address             = "disabled"
          description               = "My VRRPv3 group"
          preempt                   = "disabled"
          preempt_delay_minimum     = 10
          priority                  = 120
          vrrpv2_compatibility_mode = "enabled"
          secondary_addresses = {
            "10.1.1.11" = {}
          }
          tracked_objects = {
            "10" = {
              decrement_priority = 30
            }
          }
          vrrs_leaders = {
            "myLeader" = {}
          }
        }
      }
      vrrs_pathways = {
        "myPathway" = {
          admin_state                = "enabled"
          description                = "My pathway"
          mac_address                = "00:00:11:11:22:22"
          mac_advertisement_interval = 20
          addresses = {
            "10.1.1.20" = {}
          }
        }
      }
    }
  }
}
