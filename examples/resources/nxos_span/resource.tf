resource "nxos_span" "example" {
  sessions = {
    "1" = {
      config_state = "up"
      description  = "My SPAN session"
      type         = "local"
      source_interfaces = {
        "eth1/1" = {
          direction = "rx"
        }
      }
      source_vlans = {
        "vlan-100" = {
          direction = "rx"
        }
      }
      filter_vlans = {
        "vlan-200" = {}
      }
    }
  }
}
