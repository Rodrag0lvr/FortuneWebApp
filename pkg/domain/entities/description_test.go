package entities

import "testing"

func TestNewDescription(t *testing.T) {
	description := NewDescription(1, "Driver", "description", 1)
	if description.ID != 1 {
		t.Errorf("Expected ID to be 1 but got %d", description.ID)
	}
	if description.Input != "Driver" {
		t.Errorf("Expected Input to be Driver but got %s", description.Input)
	}
	if description.Description != "description" {
		t.Errorf("Expected Description to be description but got %s", description.Description)
	}
	if description.LiquidationId != 1 {
		t.Errorf("Expected LiquidationId to be 1 but got %d", description.LiquidationId)
	}
}

func TestNewFakeDescription(t *testing.T) {
	description := NewFakeDescription()
	if description.ID != 1 {
		t.Errorf("Expected ID to be 1 but got %d", description.ID)
	}
	if description.Input != "Driver" {
		t.Errorf("Expected Input to be Driver but got %s", description.Input)
	}
	if description.Description != "description" {
		t.Errorf("Expected Description to be description but got %s", description.Description)
	}
	if description.LiquidationId != 1 {
		t.Errorf("Expected LiquidationId to be 1 but got %d", description.LiquidationId)
	}
}
