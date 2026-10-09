resource "nxos_macsec" "example" {
  shutdown = false
  policies = {
    "POLICY1" = {
      allowed_peer_cipher_suite_1 = "GCM-AES-128"
      allowed_peer_cipher_suite_2 = "GCM-AES-256"
      allowed_peer_cipher_suite_3 = "GCM-AES-XPN-128"
      allowed_peer_cipher_suite_4 = "GCM-AES-XPN-256"
      cipher_suite                = "NoneOrEnforce-peer"
      confidentiality_offset      = "CONF-OFFSET-30"
      include_icv_indicator       = true
      include_sci                 = false
      key_server_priority         = 32
      lldp_bypass                 = false
      replay_window               = 1000
      sak_expiry_time             = "3600"
      security_policy             = "must-secure"
    }
  }
  interfaces = {
    "eth1/30" = {
      fallback_keychain_name = "MACSEC_FALLBACK_KEYCHAIN1"
      keychain_name          = "MACSEC_KEYCHAIN1"
      keychain_pki           = false
      policy_name            = "POLICY1"
    }
  }
}
