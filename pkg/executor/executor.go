package executor

import (
	"fmt"
	"github.com/Devopstrio/opentofu-blueprints/pkg/blueprint"
)

// ExecutionResult captures the output of OpenTofu plan/apply operations
type ExecutionResult struct {
	Status           string
	ResourcesCreated int
}

// TofuExecutor runs OpenTofu commands
type TofuExecutor struct{}

// NewExecutor creates a new TofuExecutor instance
func NewExecutor() *TofuExecutor {
	return &TofuExecutor{}
}

// ExecuteBlueprint simulates executing tofu plan and tofu apply
func (e *TofuExecutor) ExecuteBlueprint(bp *blueprint.Blueprint) (*ExecutionResult, error) {
	if bp == nil {
		return nil, fmt.Errorf("cannot execute nil blueprint")
	}

	return &ExecutionResult{
		Status:           "APPLIED_PROVISIONED",
		ResourcesCreated: 5,
	}, nil
}
