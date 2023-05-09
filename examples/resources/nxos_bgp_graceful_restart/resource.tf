resource "nxos_bgp_graceful_restart" "example" {
  asn              = "65001"
  vrf              = "default"
  restart_interval = 240
  stale_interval   = 1800
}
