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
		// принимаем данные
		data, ok := <-ch

		// проверяем не закрыт ли канал, в этом случае завершаем работу воркера и
		// декрементируем счетчик sync.WaitGroup
		if !ok {
			fmt.Printf("Worker %d stopped\n", id)
			return
		}

		// если канал не закрыт и есть данные - выводим их с указанием, какой воркер отработал
		fmt.Printf("Worker %d received data: %d\n", id, data)
	}
}

func Task4() {
	var wg sync.WaitGroup

	fmt.Print("Введите число воркеров: ")

	var numWorkers int

	// вводим количество воркеров
	_, err := fmt.Scanln(&numWorkers)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	// инкрементируем счетчик sync.WaitGroup на введенное количество воркеров
	wg.Add(numWorkers)

	// создаем канал для принятия данных воркерами
	ch := make(chan int)

	// запускаем введенное количество воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch, &wg)
	}

	// создаем канал для обработки сигнала Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// в качестве данных будем отдавать инкрементируемое значение i в бесконечном цикле ниже
	i := 0

	for {
		select {
		// отправляем значение i, инкрементируем его для следующей отправки, засыпаем на 200 миллисекунд
		default:
			ch <- i
			i++
			time.Sleep(200 * time.Millisecond)
		// ожидаем сигнала Ctrl+С, в этом случае закрываем канал для обработки данных в воркерах,
		// ожидаем завершения работы всех горутин (wg.Wait()) и выходим из цикла с помощью return
		case <-c:
			fmt.Println("Завершение программы..")
			close(ch)
			wg.Wait()
			fmt.Println("Программа остановлена")
			return
		}
	}
}
