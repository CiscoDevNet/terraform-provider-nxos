resource "nxos_bgp_instance_domain_peer" "example" {
  vrf              = "default"
  address          = "192.168.0.1"
  asn              = "65002"
  description      = "My description"
  peer_template    = "SPINE-PEERS"
  peer_type        = "fabric-internal"
  source_interface = "lo0"
}
