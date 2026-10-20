resource "nxos_spanning_tree" "example" {
  admin_state              = "enabled"
  instance_admin_state     = "enabled"
  bridge_assurance         = "disabled"
  control                  = "normal,stateful-ha"
  fcoe                     = "enabled"
  l2_gateway_stp_domain_id = 512
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
      mst_instances = {
        "1" = {
          cost     = 100
          priority = 64
        }
      }
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
  mst_forward_time  = 20
  mst_hello_time    = 5
  mst_max_age       = 25
  mst_maximum_hops  = 30
  mst_region_name   = "REGION1"
  mst_revision      = 5
  mst_simulate_pvst = "disabled"
  mst_instances = {
    "1" = {
      config_state = "enabled"
      priority     = 4096
      root         = "secondary"
      vlan_range   = "10-20"
    }
  }
}
