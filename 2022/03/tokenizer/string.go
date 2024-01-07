package tokenizer

import "strings"

// Half returns the left and right half of a string.
func Half(text string) (string, string) {
	mid := len(text) / 2
	return text[mid:], text[:mid]
}

// FindDuplicateChars returns a map of duplicate chars.
func FindDuplicateChars(s1, s2 string) string {
	var count = make(map[string]int)

	for _, char := range s1 {
		count[string(char)]++
	}

	var sb strings.Builder

	for _, char := range s2 {
		if _, ok := count[string(char)]; ok {
			sb.WriteString(string(char))
		}
	}
	return sb.String()
}
