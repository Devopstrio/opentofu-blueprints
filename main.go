package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/Devopstrio/opentofu-blueprints/pkg/blueprint"
	"github.com/Devopstrio/opentofu-blueprints/pkg/executor"
	"github.com/Devopstrio/opentofu-blueprints/pkg/validator"
)

func main() {
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()

	fmt.Println("==================================================")
	fmt.Println(cyan("  Devopstrio OpenTofu Blueprint Engine (tofu-runner)"))
	fmt.Println("==================================================")

	dirPath := "blueprints/aws-kubernetes"
	if len(os.Args) > 1 {
		dirPath = os.Args[1]
	}

	parser := blueprint.NewParser()
	val := validator.NewValidator()
	exec := executor.NewExecutor()

	bp, err := parser.ParseDirectory(dirPath)
	if err != nil {
		fmt.Printf("Error parsing blueprint: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[1/3] Parsed Blueprint: %s (Provider: %s)\n", bp.Name, bp.Provider)

	valRes, err := val.ValidateBlueprint(bp)
	if err != nil || !valRes.IsCompliant {
		fmt.Printf("[2/3] Policy Validation Failed!\n")
		os.Exit(1)
	}
	fmt.Printf("[2/3] Security Policy Validation: %s\n", green("COMPLIANT"))

	execRes, err := exec.ExecuteBlueprint(bp)
	if err != nil {
		fmt.Printf("[3/3] Execution Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[3/3] OpenTofu Execution Status: %s (Resources Created: %d)\n", execRes.Status, execRes.ResourcesCreated)
	fmt.Println("--------------------------------------------------")
	fmt.Println(green("SUCCESS: OpenTofu Infrastructure Blueprint Provisioned!"))
}
