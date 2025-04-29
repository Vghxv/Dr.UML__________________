package utils

import (
	"path/filepath"
	"strings"
)

// IsValidFilePath checks if a file path is valid.
// The function dislikes any invalid chars from all platforms.
// That is, even on a linux system, it will not allow the path to contain invalid chars from Windows.
func IsValidFilePath(path string) bool {
	if path == "" {
		return false
	}

	path = filepath.Clean(path)

	if strings.ContainsAny(path, `<>"|?*`) {
		return false
	}
	reserved := []string{
		"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
	}

	// 取得基本檔名（不含路徑）
	base := filepath.Base(path)
	baseName := strings.ToUpper(base)
	// 移除可能的副檔名
	if idx := strings.Index(baseName, "."); idx >= 0 {
		baseName = baseName[:idx]
	}

	for _, r := range reserved {
		if baseName == r {
			return false
		}
	}
	if strings.ContainsRune(path, '\000') {
		return false
	}

	// 檢查路徑長度
	if len(path) > 255 {
		return false
	}

	return true
}
