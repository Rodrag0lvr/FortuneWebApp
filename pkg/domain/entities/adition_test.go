package entities

import "testing"

func TestAdition(t *testing.T) {
	adition := NewAdition(1, "description", 1, 1)
	if adition.ID != 1 {
		t.Errorf("Expected ID to be 1 but got %d", adition.ID)
	}
	if adition.Description != "description" {
		t.Errorf("Expected Description to be description but got %s", adition.Description)
	}
	if adition.Price != 1 {
		t.Errorf("Expected Price to be 1 but got %f", adition.Price)
	}
	if adition.LiquidationId != 1 {
		t.Errorf("Expected LiquidationId to be 1 but got %d", adition.LiquidationId)
	}
}

func TestNewFakeAdition(t *testing.T) {
	adition := NewFakeAdition()
	if adition.ID != 1 {
		t.Errorf("Expected ID to be 1 but got %d", adition.ID)
	}
	if adition.Description != "description" {
		t.Errorf("Expected Description to be description but got %s", adition.Description)
	}
	if adition.Price != 1 {
		t.Errorf("Expected Price to be 1 but got %f", adition.Price)
	}
	if adition.LiquidationId != 1 {
		t.Errorf("Expected LiquidationId to be 1 but got %d", adition.LiquidationId)
	}
}
