resource "nxos_ipv6_access_list" "example" {
  access_lists = {
    "ACL1" = {
      extension_header   = "disabled"
      fragments          = "permit-all"
      ignore_routable    = false
      per_ace_statistics = "off"
      entries = {
        "10" = {
          ack                       = false
          action                    = "permit"
          dscp                      = 46
          destination_port_1        = "443"
          destination_port_2        = "0"
          destination_port_operator = "eq"
          destination_prefix        = "2001:db8:1::0"
          destination_prefix_length = "64"
          established               = false
          fin                       = false
          flow_label                = 0
          fragment                  = false
          http_option_type          = "invalid"
          icmp_code                 = 0
          icmp_type                 = 0
          log                       = false
          protocol                  = "tcp"
          psh                       = false
          rev                       = false
          rst                       = false
          source_port_1             = "443"
          source_port_2             = "0"
          source_port_operator      = "eq"
          source_prefix             = "2001:db8:2::0"
          source_prefix_length      = "48"
          syn                       = false
          urg                       = false
          capture_session           = 1
          dscp_mask                 = 16
          load_share                = false
          priority_all              = false
          tcp_flags_mask            = 0
          tcp_option_length         = 0
          telemetry_path            = false
          telemetry_queue           = false
        }
      }
    }
  }
  ingress_interfaces = {
    "eth1/10" = {
      access_list_name = "ACL1"
    }
  }
  ingress_vty_access_list_name = "ACL1"
  egress_interfaces = {
    "eth1/10" = {
      access_list_name = "ACL1"
    }
  }
  egress_vty_access_list_name = "ACL1"
}
