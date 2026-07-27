terraform {
  required_version = ">= 1.6.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_vpc" "tofu_vpc" {
  cidr_block           = var.vpc_cidr
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name        = "opentofu-kubernetes-vpc"
    Environment = var.environment
    ManagedBy   = "Devopstrio-OpenTofu-Blueprints"
  }
}

resource "aws_subnet" "tofu_subnet" {
  vpc_id            = aws_vpc.tofu_vpc.id
  cidr_block        = "10.10.1.0/24"
  availability_zone = "${var.aws_region}a"

  tags = {
    Name = "opentofu-private-subnet"
  }
}
