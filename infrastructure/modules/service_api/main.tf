resource "google_project_service" "cloud_run_service" {
  project = var.project_id
  service = "run.googleapis.com"
}

resource "google_project_service" "firestore_service" {
  project = var.project_id
  service = "firestore.googleapis.com"
}

resource "google_project_service" "resource_manager" {
  project = var.project_id
  service = "cloudresourcemanager.googleapis.com"
}

resource "google_project_service" "iam_credentials" {
  project = var.project_id
  service = "iamcredentials.googleapis.com"
}

resource "google_project_service" "security_token_service" {
  project = var.project_id
  service = "sts.googleapis.com"
}

# For exporting Terraform
resource "google_project_service" "cloud_asset_service" {
  project = var.project_id
  service = "cloudasset.googleapis.com"
}