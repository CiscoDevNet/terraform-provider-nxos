resource "nxos_nxapi" "example" {
  vrf                               = "management"
  http_port                         = 80
  https_port                        = 443
  idle_timeout                      = 10
  certificate_enable                = false
  certificate_file                  = "bootflash:server.crt"
  key_file                          = "bootflash:server.key"
  encrypted_key_passphrase          = "mypassphrase"
  trustpoint                        = "mytrustpoint"
  ssl_protocols                     = "TLSv1.2"
  ssl_ciphers_weak                  = false
  client_certificate_authentication = "optional"
  sudi                              = false
}
