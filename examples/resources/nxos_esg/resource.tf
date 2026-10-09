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
      selector_external_subnets_ipv4 = {
        "default;20.0.0.0/8" = {
          route_inject = "enabled"
        }
      }
      selector_external_subnets_ipv6 = {
        "default;2001:db9::/32" = {
          route_inject = "enabled"
        }
      }
      selector_external_subnets_nexthop_ipv4 = {
        "default;21.0.0.0/8;100.0.1.1" = {}
      }
      selector_external_subnets_nexthop_ipv6 = {
        "default;2001:dba::/32;2001::1" = {}
      }
      selector_external_subnets_nexthop_encap_ipv4 = {
        "default;22.0.0.0/8;100.0.1.1;mpls-vpn" = {}
      }
      selector_external_subnets_nexthop_encap_ipv6 = {
        "default;2001:dbb::/32;100.0.1.1;mpls-vpn" = {}
      }
      selector_match_interfaces = {
        "vlan100" = {}
      }
      selector_match_vlan_macs = {
        "vlan-101;00:1B:63:84:45:E6" = {}
      }
      selector_match_vlan_interfaces = {
        "vlan-102;eth1/1" = {}
      }
      type_layer4_7 = "enabled"
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
