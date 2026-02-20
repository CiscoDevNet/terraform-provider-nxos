resource "nxos_logging" "example" {
  all   = "unspecified"
  level = "information"
  facilities = [{
    name  = "spanning-tree"
    level = "information"
  }]
}
