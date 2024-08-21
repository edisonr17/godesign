package design

import (
	"testing"
)

func TestAnonymousFunction(t *testing.T) {
	tests := []struct {
		num1, num2 int16
		operation  string
		expected   int16
		expectErr  bool
	}{
		{10, 5, "+", 15, false},
		{10, 5, "-", 5, false},
		{10, 5, "*", 50, false},
		{10, 5, "/", 2, false},
		{10, 0, "/", 0, true},  // División por cero
		{10, 5, "%", 0, false}, // Operación no válida
	}

	for _, tt := range tests {
		result, err := AnonymousFunction(tt.num1, tt.num2, tt.operation)

		if (err != nil) != tt.expectErr {
			t.Errorf("AnonymousFunction(%d, %d, %q) expected error: %v, got: %v", tt.num1, tt.num2, tt.operation, tt.expectErr, err)
		}

		if result != tt.expected && !tt.expectErr {
			t.Errorf("AnonymousFunction(%d, %d, %q) = %d, expected %d", tt.num1, tt.num2, tt.operation, result, tt.expected)
		}
	}
}
