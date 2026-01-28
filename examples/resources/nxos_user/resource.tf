resource "nxos_user" "example" {
  name                     = "user1"
  allow_expired            = "yes"
  password                 = "password123"
  password_encryption_type = "clear"
  domain_name              = "all"
}
