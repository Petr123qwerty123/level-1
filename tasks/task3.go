package tasks

import (
	"fmt"
	"sync"
)

func Task3() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0

	var wg sync.WaitGroup

	wg.Add(len(numbers))

	for _, number := range numbers {
		go func(x int) {
			defer wg.Done()

			squaredX := x * x
			sum += squaredX
		}(number)
	}

	wg.Wait()
	fmt.Println(sum)
}
