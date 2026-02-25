resource "nxos_ipv6" "example" {
  vrfs = [{
    name = "VRF1"
    static_routes = [{
      prefix = "2001:db8:3333:4444:5555:6666:102:304/128"
      next_hops = [{
        interface_id = "unspecified"
        address      = "a:b::c:d/128"
        vrf_name     = "default"
        description  = "My Description"
        object       = 10
        preference   = 123
        tag          = 10
      }]
    }]
    interfaces = [{
      interface_id           = "eth1/10"
      auto_configuration     = "disabled"
      default_route          = "disabled"
      forward                = "disabled"
      link_address_use_bia   = "disabled"
      use_link_local_address = "disabled"
      urpf                   = "disabled"
      link_local_address     = "2001:db8:3333:4444:5555:6666:7777:8888"
      addresses = [{
        address = "2001:db8:3333:4444:5555:6666:7777:8888"
        type    = "primary"
        tag     = 1234
      }]
    }]
  }]
}
