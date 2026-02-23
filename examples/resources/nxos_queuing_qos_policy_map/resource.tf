resource "nxos_queuing_qos_policy_map" "example" {
  name       = "PM1"
  match_type = "match-any"
  match_class_maps = [{
    name                = "c-out-q1"
    priority            = 1
    remaining_bandwidth = 10
  }]
}
