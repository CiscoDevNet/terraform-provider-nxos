resource "nxos_spanning_tree" "example" {
  admin_state              = "enabled"
  instance_admin_state     = "enabled"
  bridge_assurance         = "disabled"
  control                  = "normal,stateful-ha"
  fcoe                     = "enabled"
  l2_gateway_stp_domain_id = 2048
  linecard_issu            = "auto"
  loopguard                = "enabled"
  mode                     = "pvrst"
  pathcost_option          = "long"
  interfaces = {
    "eth1/9" = {
      bpdu_filter               = "enable"
      bpdu_guard                = "enable"
      cost                      = 100
      guard                     = "root"
      link_type                 = "p2p"
      mode                      = "edge"
      priority                  = 200
      control                   = "bpdu-guard"
      description               = "My interface description"
      linecard_issu             = "auto"
      prestandard_configuration = "enabled"
      simulate_pvst             = "enabled"
    }
  }
  vlans = {
    "100" = {
      admin_state  = "disabled"
      diameter     = 3
      forward_time = 20
      hello_time   = 5
      max_age      = 25
      priority     = "4096"
    }
  }
}
