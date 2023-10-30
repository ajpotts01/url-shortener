module "service_account" {
  source     = "../../service_account"
  project_id = var.project_id
  env        = "dev"
  app_name   = var.app_name
}