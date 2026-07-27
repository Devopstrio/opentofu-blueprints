package executor

import (
	"testing"
	"github.com/Devopstrio/opentofu-blueprints/pkg/blueprint"
)

func TestExecuteBlueprintSuccess(t *testing.T) {
	exec := NewExecutor()
	bp := &blueprint.Blueprint{Name: "aws-kubernetes"}

	res, err := exec.ExecuteBlueprint(bp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.Status != "APPLIED_PROVISIONED" {
		t.Errorf("expected status 'APPLIED_PROVISIONED', got %s", res.Status)
	}
}

func TestExecuteBlueprintNil(t *testing.T) {
	exec := NewExecutor()
	_, err := exec.ExecuteBlueprint(nil)
	if err == nil {
		t.Error("expected error for nil blueprint, got nil")
	}
}
