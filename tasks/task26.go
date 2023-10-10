package tasks

import "strings"

// Task26 функция основана на хранении символов в ключах мапы, если по ключу уже есть значение в процессе обхода символов,
// возвращаем false, иначе true
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
