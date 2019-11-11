package mbstring

import "testing"

func TestCoalesce(t *testing.T) {
	want := "want"

	s := []string{" ", "", want, "want not"}
	result := Coalesce(s...)
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}

	s = []string{want, "want not"}
	result = Coalesce(s...)
	if result != want {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, want)
	}
}
