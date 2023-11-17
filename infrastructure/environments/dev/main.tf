module "service_api" {
  source     = "../../modules/service_api"
  app_name   = var.app_name
  project_id = var.project_id
  env        = var.env
}

module "service_account" {
  source     = "../../modules/service_account"
  project_id = var.project_id
  env        = "dev"
  app_name   = var.app_name
  depends_on = [module.service_api]
}

module "workload_identity" {
  source               = "../../modules/workload_identity"
  project_id           = var.project_id
  env                  = "dev"
  app_name             = var.app_name
  sa_provisioner_name  = var.sa_provisioner_name
  github_repo_id       = var.github_repo_id
  github_repo_owner_id = var.github_repo_owner_id
  depends_on           = [module.service_account, module.service_api]
}