package attribute

import (
	"Dr.uml/backend/component/drawdata"
	"Dr.uml/backend/utils/duerror"
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
		{
			name:      "empty content",
			attribute: Attribute{content: ""},
			expected:  "",
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
		expected  Textstyle
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
		setValue Textstyle
		hasError bool
		expected Textstyle
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
		{
			name:     "combined style",
			setValue: Bold | Italic,
			hasError: false,
			expected: Bold | Italic,
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
		initStyle Textstyle
		setValue  bool
		expected  Textstyle
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
		initStyle Textstyle
		setValue  bool
		expected  Textstyle
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
		initStyle Textstyle
		setValue  bool
		expected  Textstyle
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
		{
			name: "copy attribute with all styles",
			attribute: Attribute{
				content: "test content",
				size:    15,
				style:   Bold | Italic | Underline,
			},
			expected: Attribute{
				content: "test content",
				size:    15,
				style:   Bold | Italic | Underline,
			},
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

func TestAttribute_GetDrawData(t *testing.T) {
	tests := []struct {
		name      string
		attribute Attribute
		expected  drawdata.Attribute
		hasError  bool
	}{
		{
			name: "get draw data",
			attribute: Attribute{
				drawData: drawdata.Attribute{
					Content:   "test content",
					Height:    10,
					Width:     20,
					FontSize:  12,
					FontStyle: 3,
					FontFile:  "test.ttf",
				},
			},
			expected: drawdata.Attribute{
				Content:   "test content",
				Height:    10,
				Width:     20,
				FontSize:  12,
				FontStyle: 3,
				FontFile:  "test.ttf",
			},
			hasError: false,
		},
		{
			name:      "get empty draw data",
			attribute: Attribute{},
			expected:  drawdata.Attribute{},
			hasError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.attribute.GetDrawData()
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}
			if result.Content != tt.expected.Content {
				t.Errorf("Content: expected %v, got %v", tt.expected.Content, result.Content)
			}
			if result.Height != tt.expected.Height {
				t.Errorf("Height: expected %v, got %v", tt.expected.Height, result.Height)
			}
			if result.Width != tt.expected.Width {
				t.Errorf("Width: expected %v, got %v", tt.expected.Width, result.Width)
			}
			if result.FontSize != tt.expected.FontSize {
				t.Errorf("FontSize: expected %v, got %v", tt.expected.FontSize, result.FontSize)
			}
			if result.FontStyle != tt.expected.FontStyle {
				t.Errorf("FontStyle: expected %v, got %v", tt.expected.FontStyle, result.FontStyle)
			}
			if result.FontFile != tt.expected.FontFile {
				t.Errorf("FontFile: expected %v, got %v", tt.expected.FontFile, result.FontFile)
			}
		})
	}
}

func TestAttribute_RegisterUpdateParentDraw(t *testing.T) {
	tests := []struct {
		name     string
		hasError bool
	}{
		{
			name:     "register update function",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var att Attribute
			updateCalled := false
			updateFunc := func() duerror.DUError {
				updateCalled = true
				return nil
			}

			err := att.RegisterUpdateParentDraw(updateFunc)
			if (err != nil) != tt.hasError {
				t.Errorf("unexpected error: %v", err)
			}

			// Verify the function was stored by calling it
			if att.updateParentDraw == nil {
				t.Errorf("updateParentDraw function was not stored")
			} else {
				err = att.updateParentDraw()
				if err != nil {
					t.Errorf("unexpected error when calling updateParentDraw: %v", err)
				}
				if !updateCalled {
					t.Errorf("updateParentDraw function was not called")
				}
			}
		})
	}
}

func TestAttribute_updateDrawData(t *testing.T) {
	tests := []struct {
		name          string
		attribute     *Attribute
		expectedError bool
		updateCalled  bool
	}{
		{
			name: "update draw data with valid attribute",
			attribute: &Attribute{
				content:  "test content",
				size:     12,
				style:    Bold | Italic,
				fontFile: "test.ttf",
				drawData: drawdata.Attribute{},
				updateParentDraw: func() duerror.DUError {
					return nil
				},
			},
			expectedError: false,
			updateCalled:  true,
		},
		{
			name:          "nil attribute",
			attribute:     nil,
			expectedError: true,
			updateCalled:  false,
		},
		{
			name: "no update parent draw function",
			attribute: &Attribute{
				content:  "test content",
				size:     12,
				style:    Bold,
				fontFile: "test.ttf",
				drawData: drawdata.Attribute{},
			},
			expectedError: false,
			updateCalled:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the utils.GetTextSize function by using TestUpdateDrawData
			// which bypasses the actual text size calculation
			if tt.attribute != nil {
				updateCalled := false
				if tt.updateCalled {
					tt.attribute.updateParentDraw = func() duerror.DUError {
						updateCalled = true
						return nil
					}
				}

				err := tt.attribute.TestUpdateDrawData(10, 20, nil)

				if (err != nil) != tt.expectedError {
					t.Errorf("unexpected error: %v", err)
				}

				if !tt.expectedError && tt.attribute != nil {
					// Check if drawData was updated correctly
					if tt.attribute.drawData.Content != tt.attribute.content {
						t.Errorf("Content not updated correctly: expected %v, got %v",
							tt.attribute.content, tt.attribute.drawData.Content)
					}
					if tt.attribute.drawData.Height != 10 {
						t.Errorf("Height not updated correctly: expected %v, got %v",
							10, tt.attribute.drawData.Height)
					}
					if tt.attribute.drawData.Width != 20 {
						t.Errorf("Width not updated correctly: expected %v, got %v",
							20, tt.attribute.drawData.Width)
					}
					if tt.attribute.drawData.FontSize != tt.attribute.size {
						t.Errorf("FontSize not updated correctly: expected %v, got %v",
							tt.attribute.size, tt.attribute.drawData.FontSize)
					}
					if tt.attribute.drawData.FontStyle != int(tt.attribute.style) {
						t.Errorf("FontStyle not updated correctly: expected %v, got %v",
							int(tt.attribute.style), tt.attribute.drawData.FontStyle)
					}
					if tt.attribute.drawData.FontFile != tt.attribute.fontFile {
						t.Errorf("FontFile not updated correctly: expected %v, got %v",
							tt.attribute.fontFile, tt.attribute.drawData.FontFile)
					}

					// Check if updateParentDraw was called if it exists
					if tt.updateCalled && !updateCalled {
						t.Errorf("updateParentDraw function was not called")
					}
				}
			} else {
				err := (&Attribute{}).TestUpdateDrawData(10, 20, nil)
				if err == nil {
					t.Errorf("expected error for nil attribute, got nil")
				}
			}
		})
	}
}

func TestAttribute_TestUpdateDrawData(t *testing.T) {
	tests := []struct {
		name          string
		attribute     *Attribute
		height        int
		width         int
		err           error
		expectedError bool
	}{
		{
			name: "valid update",
			attribute: &Attribute{
				content:  "test content",
				size:     12,
				style:    Bold,
				fontFile: "test.ttf",
			},
			height:        10,
			width:         20,
			err:           nil,
			expectedError: false,
		},
		{
			name:          "nil attribute",
			attribute:     nil,
			height:        10,
			width:         20,
			err:           nil,
			expectedError: true,
		},
		{
			name: "negative height",
			attribute: &Attribute{
				content: "test content",
			},
			height:        -5,
			width:         20,
			err:           nil,
			expectedError: true,
		},
		{
			name: "negative width",
			attribute: &Attribute{
				content: "test content",
			},
			height:        10,
			width:         -5,
			err:           nil,
			expectedError: true,
		},
		{
			name: "error from parameter",
			attribute: &Attribute{
				content: "test content",
			},
			height:        10,
			width:         20,
			err:           duerror.NewInvalidArgumentError("test error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err duerror.DUError
			if tt.attribute != nil {
				err = tt.attribute.TestUpdateDrawData(tt.height, tt.width, tt.err)
			} else {
				err = (&Attribute{}).TestUpdateDrawData(tt.height, tt.width, tt.err)
			}

			if (err != nil) != tt.expectedError {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectedError && tt.attribute != nil {
				// Check if drawData was updated correctly
				if tt.attribute.drawData.Content != tt.attribute.content {
					t.Errorf("Content not updated correctly: expected %v, got %v",
						tt.attribute.content, tt.attribute.drawData.Content)
				}
				if tt.attribute.drawData.Height != tt.height {
					t.Errorf("Height not updated correctly: expected %v, got %v",
						tt.height, tt.attribute.drawData.Height)
				}
				if tt.attribute.drawData.Width != tt.width {
					t.Errorf("Width not updated correctly: expected %v, got %v",
						tt.width, tt.attribute.drawData.Width)
				}
				if tt.attribute.drawData.FontSize != tt.attribute.size {
					t.Errorf("FontSize not updated correctly: expected %v, got %v",
						tt.attribute.size, tt.attribute.drawData.FontSize)
				}
				if tt.attribute.drawData.FontStyle != int(tt.attribute.style) {
					t.Errorf("FontStyle not updated correctly: expected %v, got %v",
						int(tt.attribute.style), tt.attribute.drawData.FontStyle)
				}
				if tt.attribute.drawData.FontFile != tt.attribute.fontFile {
					t.Errorf("FontFile not updated correctly: expected %v, got %v",
						tt.attribute.fontFile, tt.attribute.drawData.FontFile)
				}
			}
		})
	}
}
