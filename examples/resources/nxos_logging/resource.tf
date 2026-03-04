resource "nxos_logging" "example" {
  all   = "unspecified"
  level = "information"
  facilities = {
    "spanning-tree" = {
      level = "information"
    }
  }
}
