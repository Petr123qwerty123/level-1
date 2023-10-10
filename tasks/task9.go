package tasks

import "fmt"

// producer - функция, принимающая на вход созданные данные и канал read в режиме записи в канал
// предназначена для записи чисел из numbers в канал read
func producer(numbers []int, read chan<- int) {
	// defer отработает в тот момент, когда все числа пройдут через read (завершится цикл)
	defer close(read)
	for _, num := range numbers {
		read <- num
	}
}

// doubler - функция, принимающая на вход канал read в режиме чтения из канала, канал write в режиме записи в канал
// предназначена для чтения из канала read, умножения прочитанного числа из канала на 2, записи в канал write
func doubler(read <-chan int, write chan<- int) {
	// defer отработает в тот момент, когда все числа пройдут через write (завершится цикл)
	defer close(write)
	for r := range read {
		double := r * 2
		write <- double
	}
}

// consumer - функция, принимающая на вход канал write в режиме чтения из канала
// предназначена для чтения из канала write удвоенных чисел и выводу в консоль
func consumer(write <-chan int) {
	for w := range write {
		fmt.Println(w)
	}
}

func Task9() {
	// создание канала для записи туда входных данных и их передачи для дальнейшей обработки
	read := make(chan int)
	// создания канала для записи обработанных данных и их передачи для вывода
	write := make(chan int)

	// заполнение входных данных
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
