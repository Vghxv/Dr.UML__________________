package utils

import (
	"testing"
)

func TestToHexString(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{
			name:     "Black",
			color:    Color{R: 0, G: 0, B: 0},
			expected: "#000000",
		},
		{
			name:     "White",
			color:    Color{R: 255, G: 255, B: 255},
			expected: "#FFFFFF",
		},
		{
			name:     "Red",
			color:    Color{R: 255, G: 0, B: 0},
			expected: "#FF0000",
		},
		{
			name:     "Green",
			color:    Color{R: 0, G: 255, B: 0},
			expected: "#00FF00",
		},
		{
			name:     "Blue",
			color:    Color{R: 0, G: 0, B: 255},
			expected: "#0000FF",
		},
		{
			name:     "Purple",
			color:    Color{R: 128, G: 0, B: 128},
			expected: "#800080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.color.ToHexString()
			if result != tt.expected {
				t.Errorf("ToHexString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFromHexString(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		expected Color
	}{
		{
			name:     "Black",
			hex:      "#000000",
			expected: Color{R: 0, G: 0, B: 0},
		},
		{
			name:     "White",
			hex:      "#FFFFFF",
			expected: Color{R: 255, G: 255, B: 255},
		},
		{
			name:     "Red",
			hex:      "#FF0000",
			expected: Color{R: 255, G: 0, B: 0},
		},
		{
			name:     "Green",
			hex:      "#00FF00",
			expected: Color{R: 0, G: 255, B: 0},
		},
		{
			name:     "Blue",
			hex:      "#0000FF",
			expected: Color{R: 0, G: 0, B: 255},
		},
		{
			name:     "Purple",
			hex:      "#800080",
			expected: Color{R: 128, G: 0, B: 128},
		},
		{
			name:     "Lowercase hex",
			hex:      "#ff00ff",
			expected: Color{R: 255, G: 0, B: 255},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromHexString(tt.hex)
			if result.R != tt.expected.R || result.G != tt.expected.G || result.B != tt.expected.B {
				t.Errorf("FromHexString(%s) = {%d, %d, %d}, want {%d, %d, %d}",
					tt.hex, result.R, result.G, result.B, tt.expected.R, tt.expected.G, tt.expected.B)
			}
		})
	}
}

func TestFromHexStringInvalidInput(t *testing.T) {
	tests := []struct {
		name string
		hex  string
	}{
		{
			name: "Empty string",
			hex:  "",
		},
		{
			name: "Too short",
			hex:  "#12345",
		},
		{
			name: "Too long",
			hex:  "#1234567",
		},
		{
			name: "Missing hash",
			hex:  "FFFFFF",
		},
		{
			name: "Invalid characters",
			hex:  "#FGHJKL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromHexString(tt.hex)
			expected := Color{} // Should return zero value for invalid inputs
			if result.R != expected.R || result.G != expected.G || result.B != expected.B {
				t.Errorf("FromHexString(%s) = {%d, %d, %d}, want {%d, %d, %d}",
					tt.hex, result.R, result.G, result.B, expected.R, expected.G, expected.B)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	colors := []Color{
		{R: 0, G: 0, B: 0},
		{R: 255, G: 255, B: 255},
		{R: 255, G: 0, B: 0},
		{R: 0, G: 255, B: 0},
		{R: 0, G: 0, B: 255},
		{R: 128, G: 128, B: 128},
		{R: 64, G: 128, B: 192},
	}

	for _, color := range colors {
		hex := color.ToHexString()
		result := FromHexString(hex)

		if result.R != color.R || result.G != color.G || result.B != color.B {
			t.Errorf("Round trip failed: original {%d, %d, %d}, got {%d, %d, %d} via %s",
				color.R, color.G, color.B, result.R, result.G, result.B, hex)
		}
	}
}
