resource "nxos_tacacs_provider_group" "example" {
  name             = "group1"
  source_interface = "lo0"
  vrf              = "management"
  deadtime         = 10
  providers = [{
    name  = "1.2.3.4"
    order = 1
  }]
}
