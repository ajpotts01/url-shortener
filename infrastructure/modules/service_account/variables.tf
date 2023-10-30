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