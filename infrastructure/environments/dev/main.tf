module "service_api" {
  source   = "../../modules/service_api"
  app_name = var.app_name
  project_id = var.project_id
  env = var.env
}

module "service_account" {
  source     = "../../modules/service_account"
  project_id = var.project_id
  env        = "dev"
  app_name   = var.app_name
  depends_on = [module.service_api]
}