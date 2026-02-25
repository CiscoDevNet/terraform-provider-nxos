resource "nxos_route_policy" "example" {
  ipv4_prefix_lists = [{
    name = "PREFIX_LIST1"
    prefix_list_entries = [{
      order      = 10
      action     = "permit"
      criteria   = "inexact"
      prefix     = "192.168.1.0/24"
      from_range = 26
      to_range   = 32
    }]
  }]
  route_maps = [{
    name = "ROUTE_MAP1"
    route_map_entries = [{
      order  = 10
      action = "permit"
      match_route_prefix_lists = [{
        prefix_list_dn = "sys/rpm/pfxlistv4-[PREFIX_LIST1]"
      }]
      additive     = "disabled"
      no_community = "disabled"
      set_criteria = "none"
      set_regular_community_items = [{
        community = "regular:as2-nn2:65001:123"
      }]
      match_tags = [{
        tag = 12345
      }]
    }]
  }]
}
