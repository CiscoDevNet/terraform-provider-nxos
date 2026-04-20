resource "nxos_keychain" "example" {
  admin_state = "enabled"
  keychains = {
    "KEYCHAIN1" = {
      keys = {
        "1" = {
          cryptographic_algorithm = "AES"
          encryption_type = "type7"
          key_string = "secret_password"
        }
      }
    }
  }
}
