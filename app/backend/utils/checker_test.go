package utils

import (
	"strings"
	"testing"
)

func TestIsValidFilePath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		hasError bool
	}{
		// Common cases
		{"EmptyPath", "", false},
		{"ValidUnixPath", "/valid/unix/path", true},
		{"ValidWindowsPath", `C:\valid\windows\path`, true},

		// Windows specific cases
		{"InvalidWindowsPathCharacters", `C:\invalid\windows|path`, false},
		{"WindowsReservedCOM1", `COM1.txt`, false},
		{"WindowsExceeds255Characters", "C:\\" + string(make([]byte, 256)), false},

		// Unix specific cases
		{"UnixContainsNullChar", "/valid/unix/path\x00", false},
		{"UnixExceeds255Characters", "/" + string(make([]byte, 256)), false},

		// Cross-platform edge cases
		{"RelativePath", "./relative/path", true},
		{"AbsolutePath", "/absolute/path", true},
		{"ValidPathWithDots", "/valid/../path", true},
		{"LongPath", strings.Repeat("a/", 128) + "file.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFilePath(tt.path)
			if (err != nil) == tt.hasError {
				t.Errorf("ValidateFilePath() error = %v, wantErr %v", err, tt.hasError)
			}
		})
	}
}
