package utils

import (
	"strings"
	"testing"

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
