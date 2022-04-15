resource "nxos_bgp_route_control" "example" {
  vrf                  = "default"
  enforce_first_as     = "disabled"
  fib_accelerate       = "enabled"
  log_neighbor_changes = "enabled"
  suppress_routes      = "disabled"
}
