package blueprint

import (
	"testing"
)

func TestParseDirectorySuccess(t *testing.T) {
	parser := NewParser()
	bp, err := parser.ParseDirectory("blueprints/aws-kubernetes")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if bp.Name != "aws-kubernetes" {
		t.Errorf("expected name 'aws-kubernetes', got %s", bp.Name)
	}

	if bp.Provider != "aws" {
		t.Errorf("expected provider 'aws', got %s", bp.Provider)
	}
}

func TestParseDirectoryEmpty(t *testing.T) {
	parser := NewParser()
	_, err := parser.ParseDirectory("")

	if err == nil {
		t.Error("expected error for empty directory path, got nil")
	}
}
