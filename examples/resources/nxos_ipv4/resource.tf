resource "nxos_ipv4" "example" {
  admin_state                                      = "enabled"
  instance_admin_state                             = "enabled"
  instance_access_list_match_local                 = "enabled"
  instance_control                                 = "stateful-ha"
  instance_hardware_ecmp_hash_offset_concatenation = "enabled"
  instance_hardware_ecmp_hash_offset_value         = 10
  instance_hardware_ecmp_hash_polynomial           = "CRC32HI"
  instance_logging_level                           = "warning"
  instance_redirect_syslog                         = "disabled"
  instance_redirect_syslog_interval                = 120
  instance_source_route                            = "disabled"
  vrfs = [{
    name                         = "VRF1"
    auto_discard                 = "enabled"
    icmp_errors_source_interface = "unspecified"
    static_routes = [{
      prefix      = "1.1.1.0/24"
      control     = "bfd"
      description = "My Description"
      preference  = 2
      tag         = 10
      next_hops = [{
        interface_id          = "unspecified"
        address               = "1.2.3.4"
        vrf_name              = "default"
        description           = "My Description"
        object                = 10
        preference            = 123
        tag                   = 10
        next_hop_name         = "nh1"
        rewrite_encapsulation = "unknown"
      }]
    }]
    interfaces = [{
      interface_id           = "eth1/10"
      drop_glean             = "disabled"
      forward                = "disabled"
      unnumbered             = "unspecified"
      urpf                   = "disabled"
      directed_broadcast_acl = "ACL1"
      directed_broadcast     = "enabled"
      addresses = [{
        address    = "24.63.46.49/30"
        type       = "primary"
        tag        = 1234
        control    = "pervasive"
        preference = 1
        use_bia    = "enabled"
        vpc_peer   = "10.0.0.1/30"
      }]
    }]
  }]
}
