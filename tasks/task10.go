package tasks

import "fmt"

func Task10() {
	step := 10
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	temperatureGroups := make(map[int][]float64)

	for _, v := range temp {
		rounded := (int(v) / step) * step
		temperatureGroups[rounded] = append(temperatureGroups[rounded], v)
	}

	fmt.Println(temperatureGroups)
}
