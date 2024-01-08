package tokenizer_test

import (
	"testing"

	"github.com/dcieplik/advent-of-code/2022/03/tokenizer"
)

func TestHalf(t *testing.T) {
	tests := map[string]struct {
		s, r1, r2 string
	}{
		"same size string":  {s: "ssss", r1: "ss", r2: "ss"},
		"different strings": {s: "ssaa", r1: "ss", r2: "aa"},
	}

	t.Parallel()
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			gotS1, gotS2 := tokenizer.Half(test.s)
			if gotS1 != test.r1 && gotS2 != test.r2 {
				t.Fatalf("got '%s' and '%s' but got '%s' and '%s'", gotS1, gotS2, test.r1, test.r2)
			}
		})
	}
}

func TestFindDuplicateChars(t *testing.T) {

	tests := map[string]struct {
		s1, s2, want string
	}{
		"no duplicates":  {s1: "abc", s2: "def", want: ""},
		"two duplicates": {s1: "abc", s2: "abs", want: "ab"},
	}
	t.Parallel()
	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := tokenizer.FindDuplicateChars(tc.s1, tc.s2)
			if got != tc.want {
				t.Fatalf("got '%s' but want '%s'", tc.want, got)
			}
		})
	}
}
