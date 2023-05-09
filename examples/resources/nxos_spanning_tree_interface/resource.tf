resource "nxos_spanning_tree_interface" "example" {
  interface_id = "eth1/9"
  admin_state = "enabled"
  bpdu_filter = "enable"
  bpdu_guard = "enable"
  cost = 100
  ctrl = "bpdu-guard"
  guard = "root"
  link_type = "p2p"
  mode = "edge"
  priority = 200
}
