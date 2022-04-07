data "nxos_bgp_instance_domain_peercontainer_af_maxprefix" "example" {
  vrf           = "default"
  template_name = "SPINE-PEERS"
  af            = "ipv4-ucast"
}
