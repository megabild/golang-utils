package mbpath

import "testing"

func TestRemoveDriveLetter(t *testing.T) {
	p := `C:\temp\folder`
	want := `\temp\folder`
	result := RemoveDriveLetter(p)
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}

	empty := ``
	resultEmpty := RemoveDriveLetter(empty)
	if resultEmpty != empty {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}
