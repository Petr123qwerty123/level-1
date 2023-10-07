package tasks

import (
	"fmt"
	"wb-level-1/set"
)

func Task12() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	s := set.NewSet()

	for _, e := range arr {
		s.Add(e)
	}

	fmt.Println(s.Items())
}
