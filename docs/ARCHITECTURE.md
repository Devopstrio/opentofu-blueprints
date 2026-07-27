# OpenTofu Infrastructure Engine Architecture

The **OpenTofu Blueprints** repository contains production-ready Infrastructure-as-Code (IaC) blueprints targeting AWS and Azure cloud platforms using **OpenTofu HCL**.

![OpenTofu Infrastructure Engine Architecture](https://raw.githubusercontent.com/Devopstrio/opentofu-blueprints/main/docs/images/architecture_diagram.jpg)

## OpenTofu CI Validation Sequence

```mermaid
flowchart TD
    GitHubActions[GitHub Actions CI Pipeline] -->|1. Checkout Repository| SetupTofu[Setup OpenTofu CLI 1.6.2]
    SetupTofu -->|2. Validate AWS Blueprint| AWSModule[blueprints/aws-kubernetes]
    SetupTofu -->|3. Validate Azure Blueprint| AzureModule[blueprints/azure-landingzone]
    AWSModule -->|tofu fmt & tofu validate| CleanAWS[AWS HCL Validated]
    AzureModule -->|tofu fmt & tofu validate| CleanAzure[Azure HCL Validated]
```

## Production OpenTofu Blueprints

1. **AWS Kubernetes Infrastructure Blueprint (`blueprints/aws-kubernetes`)**
   - High availability AWS VPC, subnets, and routing table configuration in HCL (`main.tf`, `variables.tf`, `outputs.tf`).

2. **Azure Landing Zone Infrastructure Blueprint (`blueprints/azure-landingzone`)**
   - Enterprise Azure Resource Group and Virtual Network configuration in HCL (`main.tf`, `variables.tf`, `outputs.tf`).
