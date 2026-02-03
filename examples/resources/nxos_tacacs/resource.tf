resource "nxos_tacacs" "example" {
  deadtime         = 100
  password_type    = "0"
  password         = "secret_password"
  timeout          = 10
  source_interface = "lo0"
  providers = [{
    name          = "1.2.3.4"
    auth_protocol = "chap"
    password_type = "0"
    password      = "secret_password"
    port          = 123
    retries       = 3
    timeout       = 10
  }]
}
