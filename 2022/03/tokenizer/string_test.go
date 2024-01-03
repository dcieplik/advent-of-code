package tokenizer_test

import (
	"testing"

	"github.com/dcieplik/advent-of-code/2022/03/tokenizer"
)

func TestFindDuplicateChars(t *testing.T) {

	testCases := map[string]struct {
		s1, s2, want string
	}{
		"no duplicates": {s1: "abc", s2: "def", want: ""},
	}
	for name, tc := range testCases {
		t.Parallel()
		t.Run(name, func(t *testing.T) {
			got := tokenizer.FindDuplicateChars(tc.s1, tc.s2)
			if got != tc.want {
				t.Fatalf("got '%s' but want '%s'", tc.want, got)
			}
		})
	}
}
