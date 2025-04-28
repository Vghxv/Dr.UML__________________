package Attribute

import (
	"testing"
)

func TestAttribute_GetContent(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  string
		hasError  bool
	}{
		{
			name:      "valid content",
			attribute: Attribute{content: "testContent"},
			expected:  "testContent",
			hasError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.GetContent()
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_SetContent(t *testing.T) {
	tests := []struct {
		name     string
		setValue string
		hasError bool
		expected string
	}{
		{
			name:     "valid content update",
			setValue: "newContent",
			hasError: false,
			expected: "newContent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var att Attribute
			err := att.SetContent(tt.setValue)
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if att.content != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.content)
			}
		})
	}
}

func TestAttribute_GetSize(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  int
		hasError  bool
	}{
		{
			name:      "valid size",
			attribute: Attribute{size: 10},
			expected:  10,
			hasError:  false,
		},
		{
			name:      "negative size",
			attribute: Attribute{size: -5},
			expected:  0,
			hasError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.GetSize()
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_SetSize(t *testing.T) {
	tests := []struct {
		name     string
		setValue int
		hasError bool
		expected int
	}{
		{
			name:     "valid size",
			setValue: 15,
			hasError: false,
			expected: 15,
		},
		{
			name:     "negative size",
			setValue: -5,
			hasError: true,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var att Attribute
			err := att.SetSize(tt.setValue)
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil && att.size != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.size)
			}
		})
	}
}

func TestAttribute_GetStyle(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  TextStyle
		hasError  bool
	}{
		{
			name:      "default style",
			attribute: Attribute{style: 3},
			expected:  3,
			hasError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.GetStyle()
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_SetStyle(t *testing.T) {
	tests := []struct {
		name     string
		setValue TextStyle
		hasError bool
		expected TextStyle
	}{
		{
			name:     "valid style",
			setValue: 4,
			hasError: false,
			expected: 4,
		},
		{
			name:     "invalid style",
			setValue: 8,
			hasError: true,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var att Attribute
			err := att.SetStyle(tt.setValue)
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if err == nil && att.style != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.style)
			}
		})
	}
}

func TestAttribute_SetBold(t *testing.T) {
	tests := []struct {
		name      string
		initStyle TextStyle
		setValue  bool
		expected  TextStyle
	}{
		{
			name:      "enable bold",
			initStyle: 0,
			setValue:  true,
			expected:  Bold,
		},
		{
			name:      "disable bold",
			initStyle: Bold,
			setValue:  false,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			att := Attribute{style: tt.initStyle}
			err := att.SetBold(tt.setValue)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if att.style != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.style)
			}
		})
	}
}

// SetItalic sets or unsets the italic style for the Attribute based on the value provided. Returns Utils.DUError if any error occurs.
func TestAttribute_SetItalic(t *testing.T) {
	tests := []struct {
		name      string
		initStyle TextStyle
		setValue  bool
		expected  TextStyle
	}{
		{
			name:      "enable italic",
			initStyle: 0,
			setValue:  true,
			expected:  Italic,
		},
		{
			name:      "disable italic",
			initStyle: Italic,
			setValue:  false,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			att := Attribute{style: tt.initStyle}
			err := att.SetItalic(tt.setValue)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if att.style != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.style)
			}
		})
	}
}

func TestAttribute_SetUnderline(t *testing.T) {
	tests := []struct {
		name      string
		initStyle TextStyle
		setValue  bool
		expected  TextStyle
	}{
		{
			name:      "enable underline",
			initStyle: 0,
			setValue:  true,
			expected:  Underline,
		},
		{
			name:      "disable underline",
			initStyle: Underline,
			setValue:  false,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			att := Attribute{style: tt.initStyle}
			err := att.SetUnderline(tt.setValue)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if att.style != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, att.style)
			}
		})
	}
}

func TestAttribute_IsBold(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  bool
	}{
		{
			name:      "bold style is set",
			attribute: Attribute{style: Bold},
			expected:  true,
		},
		{
			name:      "bold style is not set",
			attribute: Attribute{style: 0},
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.IsBold()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_IsItalic(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  bool
	}{
		{
			name:      "italic style is set",
			attribute: Attribute{style: Italic},
			expected:  true,
		},
		{
			name:      "italic style is not set",
			attribute: Attribute{style: 0},
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.IsItalic()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_IsUnderline(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  bool
	}{
		{
			name:      "underline style is set",
			attribute: Attribute{style: Underline},
			expected:  true,
		},
		{
			name:      "underline style is not set",
			attribute: Attribute{style: 0},
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.IsUnderline()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAttribute_Copy(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  Attribute
	}{
		{
			name: "copy attribute with all fields",
			attribute: Attribute{
				content: "test content",
				size:    10,
				style:   Bold | Italic,
			},
			expected: Attribute{
				content: "test content",
				size:    10,
				style:   Bold | Italic,
			},
		},
		{
			name:      "copy empty attribute",
			attribute: Attribute{},
			expected:  Attribute{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copy, err := tt.attribute.Copy()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if copy.content != tt.expected.content {
				t.Errorf("content: expected %v, got %v", tt.expected.content, copy.content)
			}
			if copy.size != tt.expected.size {
				t.Errorf("size: expected %v, got %v", tt.expected.size, copy.size)
			}
			if copy.style != tt.expected.style {
				t.Errorf("style: expected %v, got %v", tt.expected.style, copy.style)
			}
		})
	}
}
