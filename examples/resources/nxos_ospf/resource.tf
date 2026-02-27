resource "nxos_ospf" "example" {
  admin_state = "enabled"
  instances = [{
    name        = "OSPF1"
    admin_state = "enabled"
    control     = "stateful-ha"
    vrfs = [{
      name                          = "VRF1"
      log_adjacency_changes         = "brief"
      admin_state                   = "enabled"
      bandwidth_reference           = 400000
      bandwidth_reference_unit      = "mbps"
      distance                      = 110
      router_id                     = "34.56.78.90"
      capability_vrf_lite           = "l3vpn"
      control                       = "bfd,default-passive"
      default_metric                = 1000
      default_route_nssa_pbit_clear = true
      discard_route                 = "ext,int"
      down_bit_ignore               = true
      max_ecmp                      = 16
      name_lookup_vrf               = "default"
      rfc1583_compatible            = true
      rfc1583_compatible_ios        = true
      areas = [{
        area_id              = "0.0.0.10"
        authentication_type  = "unspecified"
        cost                 = 10
        control              = "summary"
        nssa_translator_role = "always"
        segment_routing_mpls = "mpls"
        type                 = "stub"
      }]
      max_metric_await_convergence_bgp_asn = "65535"
      max_metric_control                   = "external-lsa,startup,stub,summary-lsa"
      max_metric_external_lsa              = 600
      max_metric_summary_lsa               = 600
      max_metric_startup_interval          = 300
      interfaces = [{
        interface_id                       = "eth1/10"
        admin_state                        = "enabled"
        advertise_secondaries              = false
        area                               = "0.0.0.10"
        bfd                                = "disabled"
        cost                               = 1000
        dead_interval                      = 60
        hello_interval                     = 15
        network_type                       = "p2p"
        passive                            = "enabled"
        priority                           = 10
        control                            = "mtu-ignore"
        node_flag                          = "clear"
        retransmit_interval                = 10
        transmit_delay                     = 2
        authentication_key                 = "0 mykey"
        authentication_key_id              = 1
        authentication_key_new             = "0 mykey"
        authentication_key_secure_mode     = false
        authentication_keychain            = "mykeychain"
        authentication_md5_key             = "0 mymd5key"
        authentication_md5_key_new         = "0 mymd5key"
        authentication_md5_key_secure_mode = false
        authentication_type                = "none"
      }]
    }]
  }]
}
