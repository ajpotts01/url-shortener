variable "project_id" {
  description = "Target project ID"
  type        = string
}

variable "env" {
  description = "Which environment is this for (dev/prod)?"
  type        = string
}

variable "app_name" {
  description = "Application name"
  type        = string
}

variable "sa_provisioner_name" {
  description = "Name of service account used for provisioning"
  type        = string
}