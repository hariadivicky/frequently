package appearance

import (
	"testing"
)

func TestCounterTop(t *testing.T) {
	c := NewCounter()
	data, err := readTestSeed()
	if err != nil {
		t.Fatalf("expected no error while opening test seed; got %v", err)
	}

	result := c.Top(data, 3)

	expected := map[string]int{
		"et":  310,
		"nec": 260,
		"eu":  250,
	}

	var (
		isFound bool
		match   Result
	)

	for word, appearance := range expected {
		for _, term := range result {
			if string(term.Word) == word {
				isFound = true
				match = term
				break
			}
		}

		if !isFound {
			t.Fatalf("expected %s to be in the top list", word)
		}

		if match.Appearance != appearance {
			t.Fatalf("expected %s to have %d appearance; got %d", word, appearance, match.Appearance)
		}

		isFound = false
	}
}
