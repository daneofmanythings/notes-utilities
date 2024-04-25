package title

import (
	"bytes"
	"testing"
)

func TestTitleify(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"one char", "a", "A"},
		{"two char", "a b", "A B"},
		{"a sentence", "this is a title", "This Is A Title"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var b bytes.Buffer
			titleify(tc.input, &b)
			output := b.String()
			if output != tc.expected {
				t.Fatalf("expected=%s, got=%s", tc.expected, output)
			}
		})
	}
}
