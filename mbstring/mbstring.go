package mbstring

import "strings"

// Coalesce returns the argument that's not an empty mbstring
func Coalesce(s ...string) string {
	for _, str := range s {
		if len(strings.TrimSpace(str)) > 0 {
			return str
		}
	}
	return ""
}