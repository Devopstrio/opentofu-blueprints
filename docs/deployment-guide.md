# Developer & Infrastructure Guide: OpenTofu Blueprints

This guide outlines OpenTofu CLI installation and blueprint deployment.

## 1. Prerequisites & Installation

- **OpenTofu v1.6+**

```bash
git clone https://github.com/Devopstrio/opentofu-blueprints.git
cd opentofu-blueprints
```

## 2. Deploying AWS Kubernetes Blueprint

```bash
cd blueprints/aws-kubernetes
tofu init
tofu plan
tofu apply -auto-approve
```

## 3. Deploying Azure Landing Zone Blueprint

```bash
cd blueprints/azure-landingzone
tofu init
tofu plan
tofu apply -auto-approve
```
