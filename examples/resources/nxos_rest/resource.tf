resource "nxos_rest" "l1PhysIf" {
  dn         = "sys/intf/phys-[eth1/1]"
  class_name = "l1PhysIf"
  content = {
    id   = "eth1/1"
    mode = "trunk"
  }
}

resource "nxos_rest" "ipqosCMapInst" {
  dn         = "sys/ipqos/dflt/c/name-[CM1]"
  class_name = "ipqosCMapInst"
  content = {
    name = "CM1"
  }
  children = [
    {
      rn         = "dscp-ef"
      class_name = "ipqosDscp"
      content = {
        val = "ef"
      }
    }
  ]
}
