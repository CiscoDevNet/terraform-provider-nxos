resource "nxos_route_map" "example" {
  name = "RULE1"
  entries = [{
    order  = 10
    action = "permit"
    match_route_prefix_lists = [{
      prefix_list_dn = "sys/rpm/pfxlistv4-[LIST1]"
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
}
