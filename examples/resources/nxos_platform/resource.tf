resource "nxos_platform" "example" {
  nve_interfaces = {
    "1" = {
      ipmc_index_size = 4000
      infra_vlans = {
        "100" = {
          force = "Enable"
        }
      }
    }
  }
  extended_dme_load_interval = 300
}
