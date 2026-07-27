package validator

import (
	"testing"
	"github.com/Devopstrio/opentofu-blueprints/pkg/blueprint"
)

func TestValidateBlueprintSuccess(t *testing.T) {
	val := NewValidator()
	bp := &blueprint.Blueprint{
		Name:     "aws-kubernetes",
		Provider: "aws",
	}

	res, err := val.ValidateBlueprint(bp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !res.IsCompliant {
		t.Error("expected blueprint to be compliant")
	}
}

func TestValidateBlueprintNil(t *testing.T) {
	val := NewValidator()
	_, err := val.ValidateBlueprint(nil)
	if err == nil {
		t.Error("expected error for nil blueprint, got nil")
	}
}
