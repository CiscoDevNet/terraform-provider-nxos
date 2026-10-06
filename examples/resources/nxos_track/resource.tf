resource "nxos_track" "example" {
  objects = {
    "10" = {
      delay_down              = 10
      delay_up                = 20
      interface_id            = "eth1/10"
      interface_protocol_type = "line-protocol"
      ip_route_address_family = "ipv4"
      ip_route_prefix         = "10.1.1.0/24"
      ip_route_owner          = "none"
      ip_route_state          = "reachability"
      ip_route_vrf_name       = "default"
      ip_sla_probe_id         = 1
      ip_sla_probe_state      = "reachability"
      list_type               = "weight"
      list_weight_down        = 10
      list_weight_up          = 20
      list_members = {
        "11" = {
          weight = 20
        }
      }
    }
  }
}
