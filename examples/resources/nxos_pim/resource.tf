resource "nxos_pim" "example" {
  admin_state                    = "enabled"
  instance_admin_state           = "enabled"
  control                        = "stateful-ha"
  evpn_border_leaf               = false
  extra_net                      = false
  join_prune_delay               = 200
  null_register_delay            = 1000
  null_register_number_of_routes = 500
  register_stop                  = false
  vrfs = [{
    name                 = "default"
    admin_state          = "enabled"
    bfd                  = true
    auto_enable          = false
    control              = "flush-on-restart"
    flush_routes         = false
    join_prune_delay     = 200
    log_neighbor_changes = false
    mtu                  = 1600
    register_rate_limit  = 100
    rfc_strict           = false
    spt_switch_graceful  = false
    interfaces = [{
      interface_id         = "eth1/10"
      admin_state          = "enabled"
      bfd                  = "enabled"
      dr_priority          = 10
      passive              = false
      sparse_mode          = true
      border               = false
      border_router        = false
      control              = "border"
      description          = "MyDescription"
      dr_delay             = 5
      join_prune_route_map = "JP_POLICY"
      name                 = "pim-if"
      neighbor_route_map   = "NEIGH_POLICY"
      neighbor_prefix_list = "NEIGH_PFX"
      pfm_sd_boundary      = 0
      rfc_strict           = false
    }]
    ssm_policy_name              = "SSM"
    ssm_policy_description       = "SSM_Policy"
    ssm_range_group_list_1       = "232.0.0.0/8"
    ssm_range_group_list_2       = "233.0.0.0/8"
    ssm_range_group_list_3       = "0.0.0.0"
    ssm_range_group_list_4       = "0.0.0.0"
    ssm_range_prefix_list        = ""
    ssm_range_route_map          = ""
    ssm_range_none               = false
    static_rp_policy_name        = "RP"
    static_rp_policy_description = "Static_RP_Policy"
    static_rps = [{
      address = "1.2.3.4"
      group_lists = [{
        address  = "224.0.0.0/4"
        bidir    = true
        override = true
      }]
    }]
    anycast_rp_local_interface  = "eth1/10"
    anycast_rp_source_interface = "eth1/10"
    anycast_rp_description      = "Anycast_RP"
    anycast_rp_name             = "anycast-rp"
    anycast_rp_peers = [{
      address        = "10.1.1.1/32"
      rp_set_address = "20.1.1.1/32"
    }]
  }]
}
