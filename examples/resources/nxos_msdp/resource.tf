resource "nxos_msdp" "example" {
  admin_state          = "enabled"
  instance_admin_state = "enabled"
  instance_control     = "stateful-ha"
  vrfs = {
    "default" = {
      control                = "flush-routes"
      originator_interface   = "lo0"
      reconnect_interval     = 20
      source_active_interval = 100
      peers = {
        "10.1.1.2" = {
          admin_state                = "disabled"
          remote_as                  = "65001"
          description                = "My MSDP peer"
          keepalive_interval         = 30
          keepalive_timeout          = 60
          name                       = "peer1"
          source_active_limit        = 1000
          source_interface           = "lo0"
          authentication_description = "My authentication description"
          authentication_key         = "secret_key"
          authentication_keychain    = "KC1"
          authentication_name        = "auth1"
          authentication_type        = "md5"
          mesh_group_description     = "My mesh group description"
          mesh_group_name            = "MESH1"
          mesh_group_policy_name     = "mesh1"
          sa_policies = {
            "in" = {
              description = "My SA policy"
              name        = "sa-policy"
              route_map   = "SA_POLICY"
            }
          }
        }
      }
      group_limits = {
        "10.0.0.0/8" = {
          description    = "My group limit"
          maximum_groups = 100
          name           = "group-limit"
        }
      }
      redistribution_description = "My redistribution"
      redistribution_name        = "redist"
      redistribution_route_map   = "REDIST"
      redistribution_scope       = "inter"
      redistribution_source      = "rtmap"
    }
  }
}
