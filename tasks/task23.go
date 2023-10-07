package tasks

import "fmt"

func Task23(s []int, i int) ([]int, error) {
	if i >= 0 && i < len(s) {
		s[i] = s[len(s)-1]
		s[len(s)-1] = 0
		s = s[:len(s)-1]
		return s, nil
	} else {
		return []int{}, fmt.Errorf("invalid index")
	}
}
