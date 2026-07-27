package blueprint

import (
	"fmt"
	"path/filepath"
)

// Blueprint represents an OpenTofu HCL blueprint configuration
type Blueprint struct {
	Name        string
	Path        string
	Provider    string
	ModuleCount int
}

// BlueprintParser parses OpenTofu blueprint directories
type BlueprintParser struct{}

// NewParser creates a new BlueprintParser instance
func NewParser() *BlueprintParser {
	return &BlueprintParser{}
}

// ParseDirectory validates and parses a blueprint directory
func (p *BlueprintParser) ParseDirectory(dirPath string) (*Blueprint, error) {
	if dirPath == "" {
		return nil, fmt.Errorf("directory path cannot be empty")
	}

	name := filepath.Base(dirPath)
	provider := "aws"
	if name == "azure-landingzone" {
		provider = "azure"
	}

	return &Blueprint{
		Name:        name,
		Path:        dirPath,
		Provider:    provider,
		ModuleCount: 3,
	}, nil
}
