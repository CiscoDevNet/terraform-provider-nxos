resource "nxos_cli" "example" {
  cli = <<-EOT
  interface loopback123
  description configured-via-cli
  EOT
}
