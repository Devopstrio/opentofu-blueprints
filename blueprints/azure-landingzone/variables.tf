variable "location" {
  type        = string
  default     = "westeurope"
  description = "Azure region for landing zone"
}

variable "resource_group_name" {
  type        = string
  default     = "rg-opentofu-landingzone"
  description = "Resource group name"
}

variable "environment" {
  type        = string
  default     = "production"
  description = "Environment tier tag"
}
