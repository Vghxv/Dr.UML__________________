package utils

import (
	"runtime"
	"testing"
)

func TestIsValidFilePath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		// Common cases
		{"EmptyPath", "", false},
		{"ValidUnixPath", "/valid/unix/path", true},
		{"ValidWindowsPath", `C:\valid\windows\path`, !(runtime.GOOS == "windows")},

		// Windows specific cases
		{"InvalidWindowsPathCharacters", `C:\invalid\windows|path`, false},
		{"WindowsReservedCON", `C:\CON`, runtime.GOOS != "windows"},
		{"WindowsReservedCOM1", `COM1.txt`, runtime.GOOS != "windows"},
		{"WindowsExceeds255Characters", "C:\\" + string(make([]byte, 256)), false},

		// Unix specific cases
		{"UnixContainsNullChar", "/valid/unix/path\x00", runtime.GOOS == "windows"},
		{"UnixExceeds255Characters", "/" + string(make([]byte, 256)), false},

		// Cross-platform edge cases
		{"RelativePath", "./relative/path", true},
		{"AbsolutePath", "/absolute/path", true},
		{"ValidPathWithDots", "/valid/../path", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidFilePath(tt.path)
			if result != tt.expected {
				t.Errorf("IsValidFilePath(%q) = %v, expected %v", tt.path, result, tt.expected)
			}
		})
	}
}
