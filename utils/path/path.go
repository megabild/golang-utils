package path

import "strings"

// RemoveDriveLetterFromPath removes windows drive letter from string
func RemoveDriveLetterFromPath(s string) string {
	if s[1] == ':' {
		s = s[2:]
	}
	return s
}