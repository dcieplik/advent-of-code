package tokenizer_test

import (
	"testing"

	"github.com/dcieplik/advent-of-code/2022/03/tokenizer"
)

func TestHalf(t *testing.T) {
	tests := map[string]struct {
		s, r1, r2 string
	}{
		"ss": {s: "ssss", r1: "ss", r2: "ss"},
	}

	t.Parallel()

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
