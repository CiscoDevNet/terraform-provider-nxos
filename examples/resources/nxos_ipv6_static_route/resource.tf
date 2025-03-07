resource "nxos_ipv6_static_route" "example" {
  vrf_name = "default"
  prefix   = "2001:db8:3333:4444:5555:6666:102:304/128"
  next_hops = [{
    interface_id = "unspecified"
    address      = "a:b::c:d/128"
    vrf_name     = "default"
    description  = "My Description"
    object       = 10
    preference   = 123
    tag          = 10
  }]
}
