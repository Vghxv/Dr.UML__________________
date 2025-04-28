package utils

import (
	"testing"
)

func TestPoint_Magnitude(t *testing.T) {
	tests := []struct {
		name     string
		point    Point
		expected float64
		hasError bool
	}{
		{
			name:     "correct",
			point:    Point{X: 3, Y: 4},
			expected: float64(5),
			hasError: false,
		},
	}

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := tt.point.Magnitude()
            if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
    }
} 
