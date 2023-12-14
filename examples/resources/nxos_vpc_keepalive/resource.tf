resource "nxos_vpc_keepalive" "example" {
  destination_ip                     = "192.168.1.1"
  flush_timeout                      = 3
  interval                           = 1000
  precedence_type                    = 0
  precedence_value                   = 6
  source_ip                          = "192.168.1.2"
  timeout                            = 5
  type_of_service_byte               = 0
  type_of_service_configuration_type = 0
  type_of_service_type               = 0
  type_of_service_value              = 0
  udp_port                           = 1234
  vrf                                = "management"
}
