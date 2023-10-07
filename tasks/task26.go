package tasks

import "strings"

func Task26(s string) bool {
	m := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, c := range s {
		if m[c] {
			return false
		}

		m[c] = true
	}
	return true
}