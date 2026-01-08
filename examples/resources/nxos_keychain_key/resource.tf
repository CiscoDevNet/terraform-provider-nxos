resource "nxos_keychain_key" "example" {
  keychain   = "KEYCHAIN1"
  key_id     = 1
  key_string = "secret_password"
}
