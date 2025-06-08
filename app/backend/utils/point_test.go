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

func TestPoint_MagnitudeInt(t *testing.T) {
	tests := []struct {
		name     string
		point    Point
		expected int
		hasError bool
	}{
		{
			name:     "correct",
			point:    Point{X: 3, Y: 4},
			expected: int(5),
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.point.MagnitudeInt()
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func createTwoPoints(x1, y1, x2, y2 int) [2]Point {
	return [2]Point{
		{X: x1, Y: y1},
		{X: x2, Y: y2},
	}
}

func Test_EqualPoints(t *testing.T) {
	tests := []struct {
		name     string
		args     [2]Point
		expected bool
	}{
		{
			name:     "true",
			args:     createTwoPoints(0, 1, 0, 1),
			expected: true,
		},
		{
			name:     "false",
			args:     createTwoPoints(0, 1, 2, 3),
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EqualPoints(tt.args[0], tt.args[1])
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test_AddPoints(t *testing.T) {
	tests := []struct {
		name     string
		args     [2]Point
		expected Point
	}{
		{
			name:     "true",
			args:     createTwoPoints(0, 1, 0, 1),
			expected: Point{X: 0, Y: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddPoints(tt.args[0], tt.args[1])
			if !EqualPoints(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func Test_SubPoints(t *testing.T) {
	tests := []struct {
		name     string
		args     [2]Point
		expected Point
	}{
		{
			name:     "true",
			args:     createTwoPoints(0, 1, 2, 3),
			expected: Point{X: -2, Y: -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubPoints(tt.args[0], tt.args[1])
			if !EqualPoints(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPointStringConversion(t *testing.T) {
	strings := []string{
		"0, 0",
		"1, 2",
		"-1, -2",
		"100, 200",
		"0, -1",
	}
	for _, str := range strings {
		point, err := FromString(str)
		if err != nil {
			t.Errorf("FromString(%s) returned error: %v", str, err)
			continue
		}
		if point.String() != str {
			t.Errorf("FromString(%s) = %s, want %s", str, point.String(), str)
		}

		pointStr := point.String()
		newPoint, err := FromString(pointStr)
		if err != nil {
			t.Errorf("FromString(%s) returned error: %v", pointStr, err)
			continue
		}
		if !EqualPoints(point, newPoint) {
			t.Errorf("FromString(%s) = %v, want %v", pointStr, newPoint, point)
		}
	}
}
