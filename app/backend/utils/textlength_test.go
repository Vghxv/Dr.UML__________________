package utils

import (
	"Dr.uml/backend/drawdata"
	"os"
	"testing"
)

func Test_GetTextSize(t *testing.T) {
	tests := []struct {
		name           string
		str            string
		size           int
		fontFile       string
		expectedHeight int
		expectedWidth  int
		hasError       bool
	}{
		{
			name:           "correct",
			str:            "Hello, World!",
			size:           12,
			fontFile:       os.Getenv("APP_ROOT") + drawdata.DefaultAttributeFontFile,
			expectedHeight: 15,
			expectedWidth:  58,
			hasError:       false,
		},
		{
			name:           "size invalid",
			str:            "Hello, World!",
			size:           0,
			fontFile:       os.Getenv("APP_ROOT") + drawdata.DefaultAttributeFontFile,
			expectedHeight: 0,
			expectedWidth:  0,
			hasError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			height, width, err := GetTextSize(tt.str, tt.size, tt.fontFile)
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if height != tt.expectedHeight {
				t.Errorf("expected height %v, got %v", tt.expectedHeight, height)
			}
			if width != tt.expectedWidth {
				t.Errorf("expected width %v, got %v", tt.expectedWidth, width)
			}
		})
	}
}
