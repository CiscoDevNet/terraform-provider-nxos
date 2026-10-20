resource "nxos_vrrp" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  control              = "stateful-ha"
  interfaces = {
    "vlan10" = {
      admin_state = "enabled"
      description = "My VRRP interface"
      name        = "myVrrpIf"
      groups = {
        "10" = {
          admin_state                = "enabled"
          advertisement_interval     = 5
          authentication_text        = "cisco"
          authentication_type        = "simple"
          bfd_peer_address           = "10.1.1.5"
          forwarding_lower_threshold = 10
          forwarding_upper_threshold = 100
          preempt                    = "disabled"
          priority                   = 110
          primary_address            = "10.1.1.1"
          secondary_addresses = {
            "10.1.1.3" = {}
          }
          tracked_objects = {
            "10" = {
              decrement_priority = 20
            }
          }
        }
      }
    }
  }
}
