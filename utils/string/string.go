package string

import "strings"

// Coalesce returns the argument thats not an empty string
func Coalesce(s ...string) string {
	for _, str := range s {
		if len(strings.TrimSpace(str)) > 0 {
			return str
		}
	}
	return ""
}