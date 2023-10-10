package tasks

func Task19(s string) string {
	// создаем массив символов из переданной строки
	runes := []rune(s)

	numberRunes := len(runes)

	// меняем местами симметричные элементы массива символов переданной строки
	for i := 0; i < numberRunes/2; i++ {
		swapIndex := numberRunes - i - 1
		runes[i], runes[swapIndex] = runes[swapIndex], runes[i]
	}

	// создаем строку из массива символов
	return string(runes)
}
