resource "nxos_pim" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  vrfs = [{
    name        = "default"
    admin_state = "enabled"
    bfd         = true
    interfaces = [{
      interface_id = "eth1/10"
      admin_state  = "enabled"
      bfd          = "enabled"
      dr_priority  = 10
      passive      = false
      sparse_mode  = true
    }]
    ssm_policy_name        = "SSM"
    ssm_range_group_list_1 = "232.0.0.0/8"
    ssm_range_group_list_2 = "233.0.0.0/8"
    ssm_range_group_list_3 = "0.0.0.0"
    ssm_range_group_list_4 = "0.0.0.0"
    ssm_range_prefix_list  = ""
    ssm_range_route_map    = ""
    ssm_range_none         = false
    static_rp_policy_name  = "RP"
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
    anycast_rp_peers = [{
      address        = "10.1.1.1/32"
      rp_set_address = "20.1.1.1/32"
    }]
  }]
}
