package tasks

import (
	"fmt"
	"wb-level-1/set"
)

func Task11() {
	// создаем два пустых множества, элементы которого - int
	s1 := set.NewSet[int]()
	s2 := set.NewSet[int]()

	// заполняем первое множество числами от 0 до 4
	for i := 0; i < 5; i++ {
		s1.Add(i)
	}

	// заполняем второе множество числами от 2 до 6
	for i := 2; i < 7; i++ {
		s2.Add(i)
	}

	// ищем пересечение (от 2 до 4)
	is := s1.IntersectWith(s2)

	// выводим элементы пересечения
	fmt.Println(is.Items())
}
