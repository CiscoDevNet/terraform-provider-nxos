data "nxos_bgp_instance_domain_peercontainer" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
}
