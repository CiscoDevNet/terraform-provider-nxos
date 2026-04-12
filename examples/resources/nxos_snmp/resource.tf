resource "nxos_snmp" "example" {
  admin_state                = "enabled"
  instance_admin_state       = "enabled"
  description                = "My SNMP instance"
  engine_id                  = "00:00:00:63:00:01:00:10:20:15:10:03"
  logging_level              = "warnings"
  owner_key                  = "key1"
  owner_tag                  = "tag1"
  unknown_engine_id          = "no"
  unknown_user               = "no"
  contact                    = "admin@example.com"
  system_info_description    = "My NX-OS device"
  location                   = "DC1-Room42"
  packet_size                = 8192
  tcp_session_authentication = "tcpSessAuth"
  source_interface_traps     = "eth1/1"
  local_users = {
    "user1" = {
      authentication_password = "Xe9$mK2v!pQr"
      authentication_type     = "md5"
      ipv4_acl_name           = "snmp-acl-v4"
      ipv6_acl_name           = "snmp-acl-v6"
      enforce_privacy         = false
      localized_v2_key        = false
      localized_key           = false
      privacy_password        = "Yt4&nL8w!kDs"
      privacy_type            = "des"
      password_type           = 0
      engine_id               = "00:00:00:63:00:01:00:10:20:15:10:03"
      engine_id_length        = 0
      groups = {
        "network-operator" = {}
      }
    }
  }
  hosts = {
    "10.0.0.1;1162" = {
      community_name    = "public"
      notification_type = "traps"
      security_level    = "auth"
      version           = "v3"
    }
  }
  enable_all = "yes"
  events = {
    "1" = {
      description = "Test event"
      log         = "yes"
      owner       = "admin"
      trap        = "public"
    }
  }
}
