resource "nxos_access_list" "example" {
  access_lists = {
    "ACL1" = {
      fragments          = "permit-all"
      ignore_routable    = false
      per_ace_statistics = "off"
      udf_present        = false
      entries = {
        "10" = {
          ack                       = false
          action                    = "permit"
          dscp                      = 46
          destination_port_1        = "443"
          destination_port_2        = "0"
          destination_port_operator = "eq"
          destination_prefix        = "10.1.1.0"
          destination_prefix_length = "24"
          established               = false
          fin                       = false
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
          source_prefix             = "20.1.0.0"
          source_prefix_length      = "16"
          syn                       = false
          urg                       = false
          capture_session           = 1
          dscp_mask                 = 16
          igmp_type                 = 0
          load_share                = false
          priority_all              = false
          tcp_flags_mask            = 0
          tcp_option_length         = 0
          telemetry_path            = false
          telemetry_queue           = false
          type_of_service           = 0
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
