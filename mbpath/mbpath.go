package mbpath

// RemoveDriveLetter removes windows drive letter from mbstring
func RemoveDriveLetter(s string) string {
	if len(s) > 1 && s[1] == ':' {
		s = s[2:]
	}
	return s
}
