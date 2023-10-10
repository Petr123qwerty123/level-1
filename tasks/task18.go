package tasks

import (
	"fmt"
	"sync"
	"wb-level-1/counter"
)

func Task18() {
	// создаем Counter
	c := counter.NewCounter()
	wg := &sync.WaitGroup{}

	n := 10

	// инкрементируем счет sync.WaitGroup на число горутин
	wg.Add(n)

	// запускаем 10 горутин, в каждой из 10 горутин инкрементируем значение счетчика 1000 раз
	for i := 0; i < n; i++ {
		go func(workerNum int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				c.Inc()
			}

			fmt.Printf("Goroutine #%d is done\n", workerNum)
		}(i)
	}

	// ожидаем завершения работы всех горутин
	wg.Wait()
	// выводим значение счетчика Counter (10 * 1000 = 10000)
	fmt.Printf("Result is %d", c.GetSum())
}
