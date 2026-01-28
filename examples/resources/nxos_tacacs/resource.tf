resource "nxos_tacacs" "example" {
  deadtime         = 100
  password_type    = "0"
  password         = "secret_password"
  timeout          = 10
  source_interface = "lo0"
}
