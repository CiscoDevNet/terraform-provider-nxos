resource "nxos_nve_vni_ingress_replication" "example" {
  vni      = 103100
  protocol = "bgp"
}
