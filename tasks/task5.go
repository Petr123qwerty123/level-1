package tasks

import (
	"fmt"
	"sync"
	"time"
)

func Task5() {
	fmt.Print("Введите количество секунд работы программы: ")

	var N time.Duration

	_, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int)

	timer := time.NewTimer(N * time.Second)

	go func() {
		defer wg.Done()

		for {
			num, more := <-ch
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
		case <-timer.C:
			fmt.Printf("Время истекло, завершение программы..\n")
			close(ch)
			break writingLoop
		default:
			ch <- time.Now().Second()
			time.Sleep(time.Second)
		}
	}

	wg.Wait()
	fmt.Println("Программа остановлена")
}
