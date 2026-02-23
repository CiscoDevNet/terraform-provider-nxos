resource "nxos_ospf" "example" {
  admin_state = "enabled"
  instances = [{
    name        = "OSPF1"
    admin_state = "enabled"
    vrfs = [{
      name                     = "VRF1"
      log_adjacency_changes    = "brief"
      admin_state              = "enabled"
      bandwidth_reference      = 400000
      bandwidth_reference_unit = "mbps"
      distance                 = 110
      router_id                = "34.56.78.90"
      control                  = "bfd,default-passive"
      areas = [{
        area_id             = "0.0.0.10"
        authentication_type = "none"
        cost                = 10
        type                = "stub"
      }]
      max_metric_control          = "external-lsa,startup,stub,summary-lsa"
      max_metric_external_lsa     = 600
      max_metric_summary_lsa      = 600
      max_metric_startup_interval = 300
      interfaces = [{
        interface_id                       = "eth1/10"
        advertise_secondaries              = false
        area                               = "0.0.0.10"
        bfd                                = "disabled"
        cost                               = 1000
        dead_interval                      = 60
        hello_interval                     = 15
        network_type                       = "p2p"
        passive                            = "enabled"
        priority                           = 10
        authentication_key                 = "0 mykey"
        authentication_key_id              = 1
        authentication_key_secure_mode     = false
        authentication_keychain            = "mykeychain"
        authentication_md5_key             = "0 mymd5key"
        authentication_md5_key_secure_mode = false
        authentication_type                = "none"
      }]
    }]
  }]
}
