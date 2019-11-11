package mbpath

// RemoveDriveLetter removes windows drive letter from mbstring
func RemoveDriveLetter(s string) string {
	if s[1] == ':' {
		s = s[2:]
	}
	return s
}