resource "nxos_spanning_tree" "example" {
  interfaces = [{
    interface_id = "eth1/9"
    bpdu_filter  = "enable"
    bpdu_guard   = "enable"
    cost         = 100
    guard        = "root"
    link_type    = "p2p"
    mode         = "edge"
    priority     = 200
  }]
}
