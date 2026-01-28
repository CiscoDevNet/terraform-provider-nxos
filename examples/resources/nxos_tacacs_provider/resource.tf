resource "nxos_tacacs_provider" "example" {
  name          = "1.2.3.4"
  auth_protocol = "chap"
  password_type = "inherit-from-global"
  password      = "secret_password"
  port          = 123
  retries       = 3
  timeout       = 10
}
