package util

import "strings"

// StrCoalesce returns the argument thats not an empty string
func StrCoalesce(s ...string) string {
	for _, str := range s {
		if len(strings.TrimSpace(str)) > 0 {
			return str
		}
	}
	return ""
}

// RemoveDriveLetterFromPath removes windows drive letter from string
func RemoveDriveLetterFromPath(s string) string {
	if s[1] == ':' {
		s = s[2:]
	}
	return s
}
