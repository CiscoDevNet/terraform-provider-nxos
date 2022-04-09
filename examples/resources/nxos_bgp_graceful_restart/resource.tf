resource "nxos_bgp_graceful_restart" "example" {
  vrf              = "default"
  restart_interval = 240
  stale_interval   = 1800
}
