package tasks

import (
	"fmt"
	"wb-level-1/set"
)

func Task12() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем пустое множество
	s := set.NewSet[string]()

	// заполняем множество элементами слайса arr
	for _, e := range arr {
		s.Add(e)
	}

	// выводим элементы множества
	fmt.Println(s.Items())
}
