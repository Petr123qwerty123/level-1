package tasks

import "fmt"

func Task10() {
	// шаг группировки температур
	step := 10
	// слайс температур
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	temperatureGroups := make(map[int][]float64)

	// итерируемся по слайсу температур
	for _, v := range temp {
		// определение ключа с учетом шага
		rounded := (int(v) / step) * step
		// запись температуры в слайс соответсвующий ключу с учетом шага
		temperatureGroups[rounded] = append(temperatureGroups[rounded], v)
	}

	fmt.Println(temperatureGroups)
}
