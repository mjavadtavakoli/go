package main

import (
	"testing"
)

func TestMultiplay(t *testing.T) {
	result := Multiplay(2, 3)
	expected := 6

	if result != expected {
		t.Errorf("Multiplay(2,3) = %d; expected %d", result, expected)
	}
}
