resource "nxos_nve_vni" "example" {
  vni                            = 103100
  associate_vrf                  = false
  multicast_group                = "239.1.1.1"
  multisite_ingrress_replication = "disable"
  suppress_arp                   = "off"
}
