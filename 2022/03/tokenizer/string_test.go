package tokenizer_test

import (
	"strings"
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

// FindDuplicateChars returns a map of duplicate chars.
func FindDuplicateChars2(ss []string) string {

	find := func(s string) map[string]struct{} {
		var duplicate = make(map[string]struct{})
		for _, char := range s {

			duplicate[string(char)] = struct{}{}
		}
		return duplicate
	}
	var count = make(map[string]int)

	for _, s := range ss {
		duplicate := find(s)
		for c := range duplicate {
			count[string(c)]++
		}
	}

	var sb strings.Builder
	length := len(ss)
	for key, value := range count {
		if length == value {
			sb.WriteString(string(key))
		}
	}
	return sb.String()
}

func TestFindDuplicateChars2(t *testing.T) {

	tests := map[string]struct {
		want string
		s    []string
	}{
		"'b' as a duplicate": {s: []string{"aab", "bbb"}, want: "b"},
		"no duplicate":       {s: []string{"aab", "ccc"}, want: ""},
		"'ab' as duplicate":  {s: []string{"aab", "aab"}, want: "ab"},
	}
	t.Parallel()

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := FindDuplicateChars2(tc.s)
			if got != tc.want {
				t.Fatalf("want '%s' but got '%s'", tc.want, got)
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
