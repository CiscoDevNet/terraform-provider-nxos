resource "nxos_breakout" "example" {
  modules = {
    "1" = {
      front_panel_ports = {
        "49" = {
          breakout_map = "10g-4x"
        }
      }
    }
  }
}
