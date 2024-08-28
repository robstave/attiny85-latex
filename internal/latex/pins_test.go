package latex

import (
	"testing"
)

func TestPinValue(t *testing.T) {
	tests := []struct {
		pin      int
		expected float64
		err      bool
	}{
		{1, 1.5, false},
		{2, 0.5, false},
		{3, -0.5, false},
		{4, -1.5, false},
		{5, -1.5, false},
		{6, -0.5, false},
		{7, 0.5, false},
		{8, 1.5, false},
	}

	for _, test := range tests {
		value := PinValue(test.pin)

		if value != test.expected {
			t.Errorf("PinValue(%d) = %f; want %f", test.pin, value, test.expected)
		}
	}
}
