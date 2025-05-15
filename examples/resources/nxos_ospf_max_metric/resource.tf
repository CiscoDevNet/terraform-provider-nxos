resource "nxos_ospf_max_metric" "example" {
  instance_name               = "OSPF1"
  vrf_name                    = "VRF1"
  max_metric_control          = "external-lsa,startup,stub,summary-lsa"
  max_metric_external_lsa     = 600
  max_metric_summary_lsa      = 600
  max_metric_startup_interval = 300
}
