package tasks

import "time"

// Task25 ожидаем получение сигнала от канала, на этом основана реализованная функция sleep
func Task25(d time.Duration) {
	<-time.After(d)
}
