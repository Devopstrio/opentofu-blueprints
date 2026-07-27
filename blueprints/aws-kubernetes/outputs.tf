output "vpc_id" {
  value       = aws_vpc.tofu_vpc.id
  description = "ID of the OpenTofu provisioned VPC"
}

output "subnet_id" {
  value       = aws_subnet.tofu_subnet.id
  description = "ID of the private subnet"
}
