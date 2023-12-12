resource "google_firestore_database" "firestore_database" {
  project     = var.project_id
  location_id = "australia-southeast1" # ???
  type        = "FIRESTORE_NATIVE"
  name        = "${var.project_id}-${var.app_name}-${var.env}"
}