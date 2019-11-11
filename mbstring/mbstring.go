package mbstring

import "strings"

// Coalesce returns the first argument which is not an empty string
func Coalesce(s ...string) string {
	for _, str := range s {
		if len(strings.TrimSpace(str)) > 0 {
			return str
		}
	}
	return ""
}