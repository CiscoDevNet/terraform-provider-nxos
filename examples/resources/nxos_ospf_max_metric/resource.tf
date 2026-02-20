resource "nxos_ospf_max_metric" "example" {
  instance_name    = "OSPF1"
  vrf_name         = "VRF1"
  control          = "external-lsa,startup,stub,summary-lsa"
  external_lsa     = 600
  summary_lsa      = 600
  startup_interval = 300
}
