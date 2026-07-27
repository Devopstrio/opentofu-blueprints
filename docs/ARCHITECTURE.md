# OpenTofu Infrastructure Engine Architecture

The **OpenTofu Blueprints Engine** provides high-performance, open-source IaC orchestration using native **Golang** CLI runners and **OpenTofu HCL** infrastructure modules.

![OpenTofu Infrastructure Engine Architecture](images/architecture_diagram.jpg)

## Component Sequence Diagram

```mermaid
flowchart TD
    CLI[tofu-runner Golang CLI] -->|1. Parse HCL Path| Parser[Go Blueprint Parser Module]
    Parser -->|2. Validate Zero-Trust Rules| Validator[Go Security Validator Module]
    Validator --> IsCompliant{Is Policy Compliant?}
    IsCompliant -- Violations Found --> Halt[Halt & Exit 1]
    IsCompliant -- Policy Compliant --> Executor[Go OpenTofu Executor Engine]
    Executor -->|3. Execute tofu plan & apply| MultiCloud[AWS / Azure Cloud Subsystems]
```

## Core Infrastructure & Golang Packages

1. **Golang Blueprint Parser (`pkg/blueprint/parser.go`)**
   - Parses OpenTofu HCL directories and metadata.

2. **Golang Security Validator (`pkg/validator/validator.go`)**
   - Evaluates zero-trust policy rules against blueprint resources.

3. **Golang OpenTofu Executor (`pkg/executor/executor.go`)**
   - Orchestrates OpenTofu binary execution (`tofu plan`, `tofu apply`).

4. **Native OpenTofu HCL Modules (`blueprints/`)**
   - Production HCL modules (`main.tf`, `variables.tf`, `outputs.tf`).
