variable "aws_region" {
  type        = string
  default     = "us-east-1"
  description = "AWS deployment region for OpenTofu"
}

variable "vpc_cidr" {
  type        = string
  default     = "10.10.0.0/16"
  description = "CIDR block for Kubernetes VPC"
}

variable "environment" {
  type        = string
  default     = "production"
  description = "Deployment stage tag"
}
