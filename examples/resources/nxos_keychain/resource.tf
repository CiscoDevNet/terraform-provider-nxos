resource "nxos_keychain" "example" {
  admin_state = "enabled"
  keychains = {
    "KEYCHAIN1" = {
      keys = {
        "1" = {
          cryptographic_algorithm = "AES"
          key_string              = "secret_password"
        }
      }
    }
  }
}
