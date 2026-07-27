<div align="center">

<img src="https://raw.githubusercontent.com/Devopstrio/.github/main/assets/Browser_logo.png" height="90" alt="Devopstrio Logo" />

# opentofu-blueprints

### OpenTofu HCL Blueprint Engine & Golang Infrastructure Orchestrator

[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=flat-square)](https://devopstrio.co.uk)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8.svg?style=flat-square)](https://go.dev)
[![OpenTofu Version](https://img.shields.io/badge/OpenTofu-v1.6%2B-FF5722?style=flat-square)](https://opentofu.org)

</div>

---

## ⚡ Technical Overview & Pure Golang Architecture

The **OpenTofu Blueprints Engine** is a high-performance infrastructure orchestration platform written in **pure Golang (Go)** paired with **OpenTofu HCL** infrastructure modules.

It eliminates Python dependencies for cloud native automation, leveraging Go's native speed, strict typing, and concurrency to parse OpenTofu blueprints, validate zero-trust security baselines, and execute multi-cloud infrastructure plans.

![OpenTofu Infrastructure Engine Architecture](docs/images/architecture_diagram.jpg)

---

## 🔄 Golang Execution Sequence Flow

```mermaid
flowchart TD
    CLI[tofu-runner Golang CLI] -->|1. Parse HCL Path| Parser[Go Blueprint Parser Module]
    Parser -->|2. Validate Zero-Trust Rules| Validator[Go Security Validator Module]
    Validator --> IsCompliant{Is Policy Compliant?}
    IsCompliant -- Violations Found --> Halt[Halt & Exit 1]
    IsCompliant -- Policy Compliant --> Executor[Go OpenTofu Executor Engine]
    Executor -->|3. Execute tofu plan & apply| MultiCloud[AWS / Azure Cloud Subsystems]
```

---

## 📂 Repository Directory Layout

```
opentofu-blueprints/
├── .github/
│   └── workflows/
│       └── tofu-ci.yml          # GitHub Actions Go & OpenTofu CI pipeline
├── docs/
│   ├── ARCHITECTURE.md          # Architecture specification document
│   ├── deployment-guide.md      # Integration & deployment manual
│   └── images/
│       └── architecture_diagram.jpg # Visual blueprint diagram
├── go.mod                       # Go module manifest
├── go.sum                       # Go checksum lockfile
├── main.go                      # Go CLI binary entrypoint (tofu-runner)
├── pkg/
│   ├── blueprint/
│   │   ├── parser.go            # Go HCL blueprint parser
│   │   └── parser_test.go       # Native Go unit tests
│   ├── validator/
│   │   ├── validator.go         # Go security policy validator
│   │   └── validator_test.go    # Native Go unit tests
│   └── executor/
│       ├── executor.go          # Go OpenTofu execution engine
│       └── executor_test.go     # Native Go unit tests
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
└── README.md                    # Engine manual documentation
```

---

## 🚀 Quick Start Guide

### 1. Build Golang CLI Binary

```bash
# Clone repository
git clone https://github.com/Devopstrio/opentofu-blueprints.git
cd opentofu-blueprints

# Build Go binary
go build -o tofu-runner main.go
```

### 2. Run OpenTofu Blueprint Provisioner

```bash
./tofu-runner blueprints/aws-kubernetes
```

### 3. Run Native Go Test Suite

```bash
go test -v ./...
```

<div align="center">

<sub>&copy; 2026 Devopstrio &mdash; Engineering Uninterrupted Global Workforce Productivity.</sub>

</div>
