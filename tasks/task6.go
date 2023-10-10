package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Для выполнения этого задания я мыслил следующим образом: остановить горутину можно двумя способами - извне и изнутри

Извне:
- отправка сигнала о завершении через канал
- отправка сигнала о завершении через контекст (отмена контекста, таймаут)
- отправка сигнала о завершении через глобальную переменную
- закрытие канала (проверка закрытия канала)


Изнутри:
- логическое завершение работы горутины (отсутствие бесконечного цикла или валидное завершение этого цикла данными
неполученными извне (например, с помощью time.After()))
*/

// Task6_1 отправка сигнала о завершении через канал
func Task6_1() {
	var wg sync.WaitGroup

	done := make(chan bool)

	wg.Add(1)

	go func(d <-chan bool, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case <-d:
				fmt.Println("stop go func")
				return
			default:
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}(done, &wg)

	time.Sleep(5 * time.Second)

	done <- true

	wg.Wait()
	fmt.Println("main stop")
}

// Task6_2 отправка сигнала о завершении через контекст (отмена контекста)
func Task6_2() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("go func stop")
				return
			default:
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx, &wg)

	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()
	fmt.Println("main stop")
}

// Task6_3 отправка сигнала о завершении через контекст (таймаут)
func Task6_3() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("stop go func")
				return
			default:
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx, &wg)

	wg.Wait()
	fmt.Println("main stop")
}

// Task6_4 отправка сигнала о завершении через глобальную переменную
func Task6_4() {
	var wg sync.WaitGroup

	stop := false

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			if stop {
				fmt.Println("stop go func")
				return
			}

			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	}(&wg)

	time.Sleep(5 * time.Second)

	stop = true

	wg.Wait()
	fmt.Println("main stop")
}

// Task6_5 закрытие канала (проверка закрытия канала)
func Task6_5() {
	var wg sync.WaitGroup

	done := make(chan int, 5)

	for i := 0; i < 5; i++ {
		done <- i
	}

	wg.Add(1)

	go func(d <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		for {
			_, ok := <-d

			if !ok {
				fmt.Println("stop go func")
				return
			}

			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	}(done, &wg)

	time.Sleep(5 * time.Second)

	close(done)

	wg.Wait()
	fmt.Println("main stop")
}

// Task6_6 логическое завершение работы горутины (отсутствие бесконечного цикла или валидное завершение этого цикла данными
// неполученными извне (например, с помощью time.After()))
func Task6_6() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		ta := time.After(time.Second * 5)

		for {
			select {
			case <-ta:
				fmt.Println("stop go func")
				return
			default:
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}(&wg)

	wg.Wait()
	fmt.Println("main stop")
}
