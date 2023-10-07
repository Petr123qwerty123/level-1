package tasks

import (
	"fmt"
	"sync"
)

func Task2() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, number := range numbers {
		go func(x int) {
			defer wg.Done()

			squaredX := x * x
			fmt.Println(squaredX)
		}(number)
	}

	wg.Wait()
}
