package tasks

import (
	"fmt"
	"sync"
	"wb-level-1/counter"
)

func Task18() {
	c := counter.NewCounter()
	wg := &sync.WaitGroup{}

	n := 10

	wg.Add(n)

	// запускаем 10 горутин
	for i := 0; i < n; i++ {
		go func(workerNum int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				c.Inc()
			}

			fmt.Printf("Goroutine #%d is done\n", workerNum)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Result is %d", c.GetSum())
}
