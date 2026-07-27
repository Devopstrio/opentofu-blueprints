# Developer & Infrastructure Guide: OpenTofu Blueprints

This guide outlines Golang binary compilation, OpenTofu HCL deployment, and automated testing.

## 1. Prerequisites & Installation

- **Go v1.22+**
- **OpenTofu v1.6+**

```bash
git clone https://github.com/Devopstrio/opentofu-blueprints.git
cd opentofu-blueprints

# Build Golang CLI runner
go build -o tofu-runner main.go
```

## 2. Running Golang Runner CLI

```bash
./tofu-runner blueprints/aws-kubernetes
```

## 3. Direct OpenTofu HCL Execution

```bash
cd blueprints/aws-kubernetes
tofu init
tofu plan
tofu apply -auto-approve
```

## 4. Running Native Go Unit Tests

```bash
go test -v ./...
```
