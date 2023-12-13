resource "google_artifact_registry_repository" "app_registry" {
  project       = var.project_id
  location      = "australia-southeast1"
  repository_id = "${var.app_name}-${var.env}"
  description   = "Container registry for the ${var.app_name} app"
  format        = "DOCKER"
}