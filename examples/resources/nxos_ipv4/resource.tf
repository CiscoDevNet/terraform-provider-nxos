resource "nxos_ipv4" "example" {
  admin_state = "enabled"
  instance_admin_state = "enabled"
  access_list_match_local = "enabled"
  control = "stateful-ha"
  hardware_ecmp_hash_offset_concatenation = "enabled"
  hardware_ecmp_hash_offset_value = 10
  hardware_ecmp_hash_polynomial = "CRC32HI"
  logging_level = "warning"
  redirect_syslog = "disabled"
  redirect_syslog_interval = 120
  source_route = "disabled"
  vrfs = {
    "VRF1" = {
      auto_discard = "enabled"
      icmp_errors_source_interface = "unspecified"
      static_routes = {
        "1.1.1.0/24" = {
          control = "bfd"
          description = "My Description"
          preference = 2
          tag = 10
          next_hops = {
            "unspecified;1.2.3.4;default" = {
              description = "My Description"
              object = 10
              preference = 123
              tag = 10
              name = "nh1"
              rewrite_encapsulation = "unknown"
            }
          }
        }
      }
      interfaces = {
        "eth1/10" = {
          drop_glean = "disabled"
          forward = "disabled"
          unnumbered = "unspecified"
          urpf = "disabled"
          directed_broadcast_acl = "ACL1"
          directed_broadcast = "enabled"
          addresses = {
            "24.63.46.49/30" = {
              type = "primary"
              tag = 1234
              control = "pervasive"
              preference = 1
              use_bia = "enabled"
              vpc_peer = "10.0.0.1/30"
            }
          }
        }
      }
    }
  }
}
