package gostack

import "testing"

func Check(t *testing.T, result ElementType, correct ElementType) {
	if result != correct {
		t.Errorf("result == %d, want %d", result, correct)
	}
}
