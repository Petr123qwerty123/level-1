package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task5() {
	fmt.Print("Введите количество секунд работы программы: ")

	var N time.Duration

	// вводим количество секунд работы программы
	_, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// создаем канал для обработки данных
	ch := make(chan int)

	// создаем таймер на введенное количество секунд работы программы
	timer := time.NewTimer(N * time.Second)

	go func() {
		defer wg.Done()

		for {
			// принимаем данные и проверяем закрытие канала
			num, more := <-ch

			// если канал открыт и есть данные - выводим, в ином случае останавливаем горутину с помощью return
			// декрементируем значение счетчика sync.WaitGroup
			if more {
				fmt.Printf("Полученное значение: %ds\n", num)
			} else {
				fmt.Println("Остановка горутины чтения..")
				return
			}
		}
	}()

writingLoop:
	for {
		select {
		// если пришел сигнал об истечении времени - закрываем канал для обработки данных и
		// выходим из бесконечного цикла
		case <-timer.C:
			fmt.Printf("Время истекло, завершение программы..\n")
			close(ch)
			break writingLoop
		// отсылаем в канал в качестве данных значение секунд текущего времени и засыпаем на секунду
		default:
			ch <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}

	// ожидаем завершение работы горутины
	wg.Wait()
	fmt.Println("Программа остановлена")
}
