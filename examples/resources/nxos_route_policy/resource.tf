resource "nxos_route_policy" "example" {
  admin_state = "enabled"
  ipv4_prefix_lists = [{
    name        = "PREFIX_LIST1"
    description = "My prefix list"
    mode        = "IPV4"
    entries = [{
      order      = 10
      action     = "permit"
      criteria   = "inexact"
      prefix     = "192.168.1.0/24"
      from_range = 26
      to_range   = 32
      mask       = "255.255.255.0"
    }]
  }]
  route_maps = [{
    name           = "ROUTE_MAP1"
    pbr_statistics = "enabled"
    entries = [{
      order                   = 10
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
      match_route_prefix_lists = [{
        prefix_list_dn = "sys/rpm/pfxlistv4-[PREFIX_LIST1]"
      }]
      set_regular_community_additive     = "disabled"
      set_regular_community_no_community = "disabled"
      set_regular_community_criteria     = "none"
      set_regular_community_items = [{
        community   = "regular:as2-nn2:65001:123"
        description = "My community"
        name        = "COMMUNITY1"
      }]
      match_tags = [{
        tag = 12345
      }]
    }]
  }]
}
