package utils

import (
	"strings"
	"testing"

	"Dr.uml/backend/utils/duerror"
	"github.com/stretchr/testify/assert"
)

func TestValidateFilePath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		hasError bool
	}{
		// Common cases
		{"EmptyPath", "", true},
		{"ValidUnixPath", "/valid/unix/path", false},
		{"ValidWindowsPath", `C:\valid\windows\path`, false},

		// Windows specific cases
		{"InvalidWindowsPathCharacters", `C:\invalid\windows|path`, true},
		{"WindowsReservedCOM1", `COM1.txt`, true},
		{"WindowsExceeds255Characters", "C:\\" + string(make([]byte, 256)), true},

		// Unix specific cases
		{"UnixContainsNullChar", "/valid/unix/path\x00", true},
		{"UnixExceeds255Characters", "/" + string(make([]byte, 256)), true},

		// Cross-platform edge cases
		{"RelativePath", "./relative/path", false},
		{"AbsolutePath", "/absolute/path", false},
		{"ValidPathWithDots", "/valid/../path", false},
		{"LongPath", strings.Repeat("a/", 128) + "file.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFilePath(tt.path)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateSection(t *testing.T) {
	tests := []struct {
		name        string
		section     int
		numSections int
		hasError    bool
	}{
		{"ValidSection", 0, 3, false},
		{"ValidSectionMiddle", 1, 3, false},
		{"ValidSectionLast", 2, 3, false},
		{"NegativeSection", -1, 3, true},
		{"SectionEqualToNumSections", 3, 3, true},
		{"SectionGreaterThanNumSections", 4, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSection(tt.section, tt.numSections)
			if tt.hasError {
				assert.Error(t, err)
				assert.IsType(t, duerror.NewInvalidArgumentError(""), err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateIndex(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		numItems int
		hasError bool
	}{
		{"ValidIndex", 0, 3, false},
		{"ValidIndexMiddle", 1, 3, false},
		{"ValidIndexLast", 2, 3, false},
		{"NegativeIndex", -1, 3, true},
		{"IndexEqualToNumItems", 3, 3, true},
		{"IndexGreaterThanNumItems", 4, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateIndex(tt.index, tt.numItems)
			if tt.hasError {
				assert.Error(t, err)
				assert.IsType(t, duerror.NewInvalidArgumentError(""), err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
