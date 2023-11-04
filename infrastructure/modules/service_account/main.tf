resource "google_service_account" "app_service_account" {
  account_id   = "ajp-${var.app_name}-${var.env}-sa"
  project      = var.project_id
  display_name = "Cloud Resume Service Account"
}