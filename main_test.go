package main

import (
	"testing"
	"github.com/Devopstrio/opentofu-blueprints/pkg/blueprint"
	"github.com/Devopstrio/opentofu-blueprints/pkg/validator"
	"github.com/Devopstrio/opentofu-blueprints/pkg/executor"
)

func TestMainOrchestration(t *testing.T) {
	parser := blueprint.NewParser()
	val := validator.NewValidator()
	exec := executor.NewExecutor()

	bp, err := parser.ParseDirectory("blueprints/aws-kubernetes")
	if err != nil {
		t.Fatalf("failed to parse directory: %v", err)
	}

	valRes, err := val.ValidateBlueprint(bp)
	if err != nil || !valRes.IsCompliant {
		t.Fatalf("failed validation: %v", err)
	}

	execRes, err := exec.ExecuteBlueprint(bp)
	if err != nil {
		t.Fatalf("failed execution: %v", err)
	}

	if execRes.Status != "APPLIED_PROVISIONED" {
		t.Errorf("expected status APPLIED_PROVISIONED, got %s", execRes.Status)
	}
}
