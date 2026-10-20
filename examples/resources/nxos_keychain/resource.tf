resource "nxos_keychain" "example" {
  admin_state = "enabled"
  keychains = {
    "KEYCHAIN1" = {
      keys = {
        "1" = {
          cryptographic_algorithm     = "AES"
          key_string                  = "secret_password"
          accept_lifetime_end_day     = 31
          accept_lifetime_end_month   = "Dec"
          accept_lifetime_end_time    = "23:59:59"
          accept_lifetime_end_year    = 2035
          accept_lifetime_local       = "local"
          accept_lifetime_start_day   = 1
          accept_lifetime_start_month = "Jan"
          accept_lifetime_start_time  = "00:00:00"
          accept_lifetime_start_year  = 2025
          send_lifetime_end_day       = 31
          send_lifetime_end_month     = "Dec"
          send_lifetime_end_time      = "23:59:59"
          send_lifetime_end_year      = 2035
          send_lifetime_local         = "local"
          send_lifetime_start_day     = 1
          send_lifetime_start_month   = "Jan"
          send_lifetime_start_time    = "00:00:00"
          send_lifetime_start_year    = 2025
        }
      }
    }
  }
  macsec_keychains = {
    "MACSEC_KEYCHAIN1" = {
      keys = {
        "1000" = {
          cryptographic_algorithm = "AES_256_CMAC"
          key_hex_string          = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"
          end_day                 = 31
          end_month               = "Dec"
          end_time                = "23:59:59"
          end_year                = 2035
          local                   = "local"
          start_day               = 1
          start_month             = "Jan"
          start_time              = "00:00:00"
          start_year              = 2025
        }
      }
    }
  }
  macsec_psk_no_show = "enabled"
}
