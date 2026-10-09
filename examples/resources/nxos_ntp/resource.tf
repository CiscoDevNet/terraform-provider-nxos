resource "nxos_ntp" "example" {
  admin_state          = "enabled"
  allow_control        = "enabled"
  allow_private        = "enabled"
  authentication_state = "enabled"
  logging              = "enabled"
  logging_level        = "error"
  master               = "enabled"
  master_stratum       = 4
  passive              = "enabled"
  rate_limit           = 5
  servers = {
    "1.2.3.4" = {
      vrf       = "management"
      type      = "server"
      key_id    = 10
      min_poll  = 4
      max_poll  = 6
      preferred = true
    }
  }
  authentication_keys = {
    "10" = {
      authentication_type = "md5"
      encryption_type     = "type7"
      encrypted           = true
      key                 = "secret_key"
      name                = "KEY1"
      trusted             = true
    }
  }
  source_interface        = "lo0"
  access_group_match_all  = "enabled"
  access_group_peer       = "PeerAcl"
  access_group_query_only = "QueryOnlyAcl"
  access_group_serve      = "ServeAcl"
  access_group_serve_only = "ServeOnlyAcl"
}
