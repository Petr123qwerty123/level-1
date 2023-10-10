package tasks

import (
	"fmt"
	"sync"
	"wb-level-1/concurrencyMap"
)

func Task7() {
	var wg sync.WaitGroup

	// будем запускать 10 горутин для записи
	n := 10

	wg.Add(n)

	// создаем map с возможностью конкурентной записи, ключи - int, значения - string
	cm := concurrencyMap.NewConcurrencyMap[int, string]()

	for i := 0; i < n; i++ {
		i := i

		go func() {
			defer wg.Done()

			// записываем в map, в ключ - индекс, в значение руну по индексу
			cm.Set(i, string(rune(i)))
			fmt.Printf("%d:%s\n", i, string(rune(i)))
		}()
	}

	// ожидаем завершения работы всех горутин
	wg.Wait()

	// проверяем, что количество записанных данных равно количеству запущенных горутин
	fmt.Println(len(cm.Map) == n)
}
