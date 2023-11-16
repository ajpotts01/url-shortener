resource "google_iam_workload_identity_pool" "github_actions_pool" {
  project                   = var.project_id
  workload_identity_pool_id = "github-actions"
  display_name              = "Github Actions"
  description               = "Identity pool for CI/CD pipelines via Github Actions"
}

resource "google_iam_workload_identity_pool_provider" "github_actions_provider" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_actions_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-actions"
  display_name                       = "Github Actions"

  oidc {
    allowed_audiences = [ var.allowed_audience ]
    issuer_uri = "https://token.actions.githubusercontent.com"
  }

  attribute_mapping = {
    "google.subject"                = "assertion.sub"
    "attribute.repository_id"       = "assertion.repository_id" 
    "attribute.repository_owner_id" = "assertion.repository_owner_id"
  }

  attribute_condition = "assertion.repository_id=='${var.github_repo_id}' && assertion.repository_owner_id=='${var.github_repo_owner_id}'"
}

resource "google_service_account_iam_member" "identity_federation" {
  service_account_id = "projects/${var.project_id}/serviceAccounts/${var.sa_provisioner_name}@${var.project_id}.iam.gserviceaccount.com"
  role               = "roles/iam.workloadIdentityUser"
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_actions_pool.name}/attribute.repository_id/${var.github_repo_id}"
  depends_on         = [google_iam_workload_identity_pool_provider.github_actions_provider]
}

output "google_iam_workload_identity_pool_provider_github_name" {
  description = "Workload Identity Pool Provider ID"
  value       = google_iam_workload_identity_pool_provider.github_actions_provider.name
  depends_on  = [google_iam_workload_identity_pool_provider.github_actions_provider]
}