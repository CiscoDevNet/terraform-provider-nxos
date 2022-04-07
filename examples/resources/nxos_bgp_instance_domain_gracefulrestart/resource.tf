resource "nxos_bgp_instance_domain_gracefulrestart" "example" {
  vrf              = "default"
  restart_interval = 240
  stale_interval   = 1800
}
