package tasks

import (
	"fmt"
	"sync"
)

func Task2() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	// инкрементируем счетчик sync.WaitGroup на количество горутин (по логике задачи это число равно
	// количеству чисел в массиве numbers)
	wg.Add(len(numbers))

	for _, number := range numbers {
		// конкурентная функция, которая считает квадрат x и выводит его на экран
		go func(x int, wg *sync.WaitGroup) {
			// декрементируем счетчик sync.WaitGroup по завершению работы горутины
			defer wg.Done()

			squaredX := x * x
			fmt.Println(squaredX)
		}(number, &wg)
	}

	// ожидаем, когда значение счетчика sync.WaitGroup станет равно нулю
	wg.Wait()
}
