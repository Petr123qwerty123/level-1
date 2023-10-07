package tasks

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		data, ok := <-ch

		if !ok {
			fmt.Printf("Worker %d stopped\n", id)
			return
		}

		fmt.Printf("Worker %d received data: %d\n", id, data)
	}
}

func Task4() {
	var wg sync.WaitGroup

	fmt.Print("Введите число воркеров: ")

	var numWorkers int

	_, err := fmt.Scanln(&numWorkers)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	wg.Add(numWorkers)

	ch := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch, &wg)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	i := 0

	for {
		select {
		default:
			ch <- i
			i++
			time.Sleep(200 * time.Millisecond)
		case <-c:
			fmt.Println("Завершение программы..")
			close(ch)
			wg.Wait()
			fmt.Println("Программа остановлена")
			return
		}
	}
}
