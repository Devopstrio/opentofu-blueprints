<div align="center">

<img src="https://raw.githubusercontent.com/Devopstrio/.github/main/assets/Browser_logo.png" height="90" alt="Devopstrio Logo" />

# opentofu-blueprints

### Enterprise OpenTofu Infrastructure-as-Code (IaC) Blueprints

[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=flat-square)](https://devopstrio.co.uk)
[![OpenTofu Version](https://img.shields.io/badge/OpenTofu-v1.6%2B-FF5722?style=flat-square)](https://opentofu.org)
[![HCL Standard](https://img.shields.io/badge/IaC-OpenTofu_HCL-623CE4?style=flat-square)](https://opentofu.org)

</div>

---

## ⚡ Technical Overview & OpenTofu Scope

The **OpenTofu Blueprints** repository provides production-ready Infrastructure-as-Code (IaC) blueprints targeting AWS and Azure cloud platforms using **OpenTofu HCL**.

It contains verified HCL modules for deploying multi-cloud virtual networks, subnets, and foundational infrastructure.

![OpenTofu Infrastructure Engine Architecture](https://raw.githubusercontent.com/Devopstrio/opentofu-blueprints/main/docs/images/architecture_diagram.jpg)

---

## 🔄 OpenTofu CI Validation Flow

```mermaid
flowchart TD
    GitHubActions[GitHub Actions CI Pipeline] -->|1. Checkout Repository| SetupTofu[Setup OpenTofu CLI 1.6.2]
    SetupTofu -->|2. Validate AWS Blueprint| AWSModule[blueprints/aws-kubernetes]
    SetupTofu -->|3. Validate Azure Blueprint| AzureModule[blueprints/azure-landingzone]
    AWSModule -->|tofu fmt & tofu validate| CleanAWS[AWS HCL Validated]
    AzureModule -->|tofu fmt & tofu validate| CleanAzure[Azure HCL Validated]
```

---

## 📂 Repository Directory Layout

```
opentofu-blueprints/
├── .github/
│   └── workflows/
│       └── tofu-ci.yml          # GitHub Actions OpenTofu HCL CI pipeline
├── docs/
│   ├── ARCHITECTURE.md          # Architecture specification document
│   ├── deployment-guide.md      # Integration & deployment manual
│   └── images/
│       └── architecture_diagram.jpg # Visual blueprint diagram
├── blueprints/
│   ├── aws-kubernetes/
│   │   ├── main.tf              # AWS EKS VPC HCL blueprint
│   │   ├── variables.tf         # Input variables
│   │   └── outputs.tf           # Output variables
│   └── azure-landingzone/
│       ├── main.tf              # Azure VNet HCL blueprint
│       ├── variables.tf         # Input variables
│       └── outputs.tf           # Output variables
├── scripts/
│   ├── install-opentofu.sh      # OpenTofu installer script
│   └── verify-blueprints.sh     # OpenTofu HCL verification script
├── .gitignore                   # Git ignore file
└── README.md                    # Infrastructure manual documentation
```

---

## 🚀 Quick Start Guide

### 1. Deploy AWS Kubernetes Blueprint

```bash
cd blueprints/aws-kubernetes
tofu init
tofu plan
tofu apply -auto-approve
```

### 2. Deploy Azure Landing Zone Blueprint

```bash
cd blueprints/azure-landingzone
tofu init
tofu plan
tofu apply -auto-approve
```

<div align="center">

<sub>&copy; 2026 Devopstrio &mdash; Engineering Uninterrupted Global Workforce Productivity.</sub>

</div>
