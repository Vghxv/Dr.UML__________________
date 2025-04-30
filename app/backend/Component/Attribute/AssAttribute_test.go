package attribute

import (
	"Dr.uml/backend/utils/duerror"
	"errors"
	"testing"
)

func TestAssAttribute_GetRatio(t *testing.T) {
	tests := []struct {
		name       string
		inputRatio float64
		wantRatio  float64
		wantErr    duerror.DUError
	}{
		{"valid ratio", 0.5, 0.5, nil},
		{"zero ratio", 0.0, 0.0, nil},
		{"one ratio", 1.0, 1.0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := NewAssAttribute(tt.inputRatio)
			if err != nil {
				t.Errorf("NewAssAttribute() error = %v", err)
				return
			}
			gotRatio, gotErr := attr.GetRatio()

			if gotRatio != tt.wantRatio {
				t.Errorf("GetRatio() = %v, want %v", gotRatio, tt.wantRatio)
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("GetRatio() error = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestAssAttribute_SetRatio(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		wantErr   bool
		errorText string
	}{
		{"valid ratio", 0.5, false, ""},
		{"zero ratio", 0.0, false, ""},
		{"one ratio", 1.0, false, ""},
		{"negative ratio", -0.1, true, "ratio should be between 0 and 1"},
		{"greater than one ratio", 1.1, true, "ratio should be between 0 and 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := NewAssAttribute(0.5) // Initialize with a default valid ratio
			if err != nil {
				t.Errorf("NewAssAttribute() error = %v", err)
				return
			}
			gotErr := attr.SetRatio(tt.input)

			if tt.wantErr {
				if gotErr == nil {
					t.Errorf("SetRatio() expected error but got nil")
				} else if gotErr.Error() != tt.errorText {
					t.Errorf("SetRatio() error text = %v, want %v", gotErr.Error(), tt.errorText)
				}
			} else {
				if gotErr != nil {
					t.Errorf("SetRatio() unexpected error = %v", gotErr)
				}
			}
		})
	}
}

func TestAssAttribute_GetAssDD(t *testing.T) {
	attr, _ := NewAssAttribute(0.5)
	dd := attr.GetAssDD()
	if dd != attr.assDD {
		t.Errorf("GetAssDD() = %v, want %v", dd, attr.assDD)
	}
}

func TestAssAttribute_UpdateDrawData(t *testing.T) {
	attr, _ := NewAssAttribute(0.5)
	attr.content = "test"
	attr.size = 12
	attr.style = Bold
	attr.fontFile = "test.ttf"
	attr.Ratio = 0.7

	attr.UpdateDrawData()

	if attr.assDD.Content != attr.content {
		t.Errorf("UpdateDrawData() content = %v, want %v", attr.assDD.Content, attr.content)
	}
	if attr.assDD.FontSize != attr.size {
		t.Errorf("UpdateDrawData() size = %v, want %v", attr.assDD.FontSize, attr.size)
	}
	if attr.assDD.FontStyle != int(attr.style) {
		t.Errorf("UpdateDrawData() style = %v, want %v", attr.assDD.FontStyle, int(attr.style))
	}
	if attr.assDD.FontFile != attr.fontFile {
		t.Errorf("UpdateDrawData() fontFile = %v, want %v", attr.assDD.FontFile, attr.fontFile)
	}
	if attr.assDD.Ratio != attr.Ratio {
		t.Errorf("UpdateDrawData() ratio = %v, want %v", attr.assDD.Ratio, attr.Ratio)
	}
}

func TestNewAssAttribute(t *testing.T) {
	tests := []struct {
		name       string
		inputRatio float64
		wantRatio  float64
		wantErr    bool
		errorText  string
	}{
		{"valid ratio", 0.5, 0.5, false, ""},
		{"zero ratio", 0.0, 0.0, false, ""},
		{"one ratio", 1.0, 1.0, false, ""},
		{"negative ratio", -0.1, 0, true, "ratio should be between 0 and 1"},
		{"greater than one ratio", 1.1, 0, true, "ratio should be between 0 and 1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr, err := NewAssAttribute(tt.inputRatio)
			if tt.wantErr {
				if err == nil {
					t.Errorf("NewAssAttribute() expected error but got nil")
					return
				}
				if err.Error() != tt.errorText {
					t.Errorf("NewAssAttribute() error = %v, want %v", err.Error(), tt.errorText)
				}
				return
			}
			if err != nil {
				t.Errorf("NewAssAttribute() error = %v", err)
				return
			}
			if attr.Ratio != tt.wantRatio {
				t.Errorf("NewAssAttribute() = %v, want %v", attr.Ratio, tt.wantRatio)
			}
		})
	}
}
