data "nxos_rest" "l1PhysIf" {
  dn = "sys/intf/phys-[eth1/1]"
}
