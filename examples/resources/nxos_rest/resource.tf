resource "nxos_rest" "l1PhysIf" {
  dn         = "sys/intf/phys-[eth1/1]"
  class_name = "l1PhysIf"
  content = {
    id   = "eth1/1"
    mode = "trunk"
  }
}
