resource "guance_dashboard" "main" {
  name     = var.name
  manifest = file("${path.module}/manifest.json")
}
