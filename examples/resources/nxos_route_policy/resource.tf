resource "nxos_route_policy" "example" {
  ipv4_prefix_lists = {
    "PREFIX_LIST1" = {
      description = "My prefix list"
      mode        = "IPV4"
      entries = {
        "10" = {
          action     = "permit"
          criteria   = "inexact"
          prefix     = "192.168.1.0/24"
          from_range = 26
          to_range   = 32
          mask       = "255.255.255.0"
        }
      }
    }
  }
  route_maps = {
    "ROUTE_MAP1" = {
      pbr_statistics = "enabled"
      entries = {
        "10" = {
          action                  = "permit"
          description             = "My route map entry"
          drop_on_fail_v4         = "disabled"
          drop_on_fail_v6         = "disabled"
          force_order_v4          = "disabled"
          force_order_v6          = "disabled"
          load_share_v4           = "disabled"
          load_share_v6           = "disabled"
          set_default_next_hop_v4 = "disabled"
          set_default_next_hop_v6 = "disabled"
          set_vrf_v4              = "disabled"
          set_vrf_v6              = "disabled"
          verify_availability_v4  = "disabled"
          verify_availability_v6  = "disabled"
          match_route_prefix_lists = {
            "sys/rpm/pfxlistv4-[PREFIX_LIST1]" = {}
          }
          set_regular_community_additive     = "disabled"
          set_regular_community_no_community = "disabled"
          set_regular_community_criteria     = "none"
          set_regular_community_items = {
            "regular:as2-nn2:65001:123" = {
            }
          }
          match_tags = {
            "12345" = {}
          }
          set_metric_is_bgp                = false
          set_metric                       = "100"
          set_metric_delay                 = 10
          set_metric_load                  = 10
          set_metric_mtu                   = 1500
          set_metric_reliability           = 100
          set_metric_type                  = "type-1"
          set_next_hop_v4_peer_address     = "disabled"
          set_next_hop_v4_redist_unchanged = "enabled"
          set_next_hop_v4_unchanged        = "enabled"
          set_next_hop_v6_peer_address     = "disabled"
          set_next_hop_v6_redist_unchanged = "enabled"
          set_next_hop_v6_unchanged        = "enabled"
        }
      }
    }
  }
  community_lists = {
    "COMMUNITY_LIST1" = {
      mode = "standard"
      entries = {
        "10" = {
          action = "permit"
          items = {
            "regular:as2-nn2:65001:123" = {
            }
          }
        }
      }
    }
  }
}
