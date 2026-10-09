resource "nxos_esg" "example" {
  mac_segmentation = "enabled"
  security_groups = {
    "100" = {
      name = "group1"
      selector_connected_endpoints_ipv4 = {
        "default;10.0.0.0/24" = {}
      }
      selector_connected_endpoints_ipv6 = {
        "default;2001:db8::/32" = {}
      }
      selector_match_vlans = {
        "vlan-100" = {}
      }
    }
  }
  class_maps = {
    "cmap1" = {
      description = "My class map"
      filter_entries = {
        "entry1" = {
          apply_to_fragment           = true
          arp_opcode                  = "req"
          destination_port_from       = "300"
          destination_port_to         = "400"
          ether_type                  = "ipv4"
          icmpv4_type                 = 0
          icmpv6_type                 = 0
          match_destination_port_zero = true
          match_dscp                  = "AF11"
          match_source_port_zero      = true
          protocol                    = "tcp"
          source_port_from            = "100"
          source_port_to              = "200"
          stateful                    = true
          tcp_rules                   = "ack,syn"
        }
      }
    }
  }
  policy_maps = {
    "pmap1" = {
      description = "My policy map"
      match_class_maps = {
        "cmap1" = {
          count_action      = false
          forwarding_action = "permit"
          log_action        = false
          redirect_chain    = "chain1"
        }
      }
    }
  }
  domains = {
    "VRF1" = {
      default_action        = "deny"
      policy_classifier_tag = 200
      security_mode         = "enforced"
      contracts = {
        "100;any;pmap1" = {
          direction = "uni-dir"
        }
      }
    }
  }
}
