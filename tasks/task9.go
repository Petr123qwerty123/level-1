package tasks

import "fmt"

func producer(numbers []int, read chan<- int) {
	defer close(read)
	for _, num := range numbers {
		read <- num
	}
}

func doubler(read <-chan int, write chan<- int) {
	defer close(write)
	for r := range read {
		double := r * 2
		write <- double
	}
}

func consumer(write <-chan int) {
	for w := range write {
		fmt.Println(w)
	}
}

func Task9() {
	read := make(chan int)
	write := make(chan int)

	var numbers []int
	func() {
		for i := 1; i <= 20; i++ {
			numbers = append(numbers, i)
		}
	}()

	go producer(numbers, read)
	go doubler(read, write)
	consumer(write)
}
