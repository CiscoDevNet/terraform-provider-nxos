resource "nxos_object_group" "example" {
  ipv4_address_object_groups = {
    "AG1" = {
      members = {
        "10" = {
          prefix        = "192.168.1.0"
          prefix_length = "24"
        }
      }
    }
  }
  ipv6_address_object_groups = {
    "AG6" = {
      members = {
        "10" = {
          prefix        = "2001:db8::"
          prefix_length = "64"
        }
      }
    }
  }
  port_object_groups = {
    "PG1" = {
      members = {
        "10" = {
          port_operator = "eq"
          port_1        = "443"
        }
      }
    }
  }
}
