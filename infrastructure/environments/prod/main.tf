module "service_account" {
  source     = "../../modules/service_account"
  project_id = var.project_id
  env        = var.env
  app_name   = var.app_name
}

module "firestore" {
  source     = "../../modules/firestore"
  project_id = var.project_id
  env        = "prod"
  app_name   = var.app_name
  depends_on = [module.service_account]
}