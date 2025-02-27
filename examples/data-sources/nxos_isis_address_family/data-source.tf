data "nxos_isis_address_family" "example" {
  instance_name  = "ISIS1"
  vrf            = "default"
  address_family = "v4"
}
