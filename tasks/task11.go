package tasks

import (
	"fmt"
	"wb-level-1/set"
)

func Task11() {
	s1 := set.NewSet()
	s2 := set.NewSet()

	for i := 0; i < 5; i++ {
		s1.Add(i)
	}

	for i := 2; i < 7; i++ {
		s2.Add(i)
	}

	is := s1.IntersectWith(s2)

	fmt.Println(is.Items())
}
