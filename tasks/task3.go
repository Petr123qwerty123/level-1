package tasks

import (
	"fmt"
	"sync"
)

func Task3() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0

	var wg sync.WaitGroup

	// инкрементируем счетчик sync.WaitGroup на количество горутин (по логике задачи это число равно
	// количеству чисел в массиве numbers)
	wg.Add(len(numbers))

	for _, number := range numbers {
		// конкурентная функция, которая считает квадрат x и прибавляет его к sum
		go func(x int, wg *sync.WaitGroup, sum *int) {
			// декрементируем счетчик sync.WaitGroup по завершению работы горутины
			defer wg.Done()

			squaredX := x * x
			*sum += squaredX
		}(number, &wg, &sum)
	}

	// ожидаем, когда значение счетчика sync.WaitGroup станет равно нулю
	wg.Wait()
	// выводим значение суммы квадратов
	fmt.Println(sum)
}
