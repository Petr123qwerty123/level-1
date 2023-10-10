package tasks

import (
	"regexp"
	"strings"
)

// в этой задаче словами будем считать подряд идущие непробельные символы

func Task20(s string) string {
	var reverseS string

	// создаем регулярное выражение для поиска подряд идущих непробельных символов
	reWords, _ := regexp.Compile(`(\S)+`)
	// записываем найденные слова в слайс
	words := reWords.FindAllString(s, -1)

	numberWords := len(words)

	// меняем местами симметричные элементы массива слов из переданной строки
	for i := 0; i < numberWords/2; i++ {
		swapIndex := numberWords - i - 1
		words[i], words[swapIndex] = words[swapIndex], words[i]
	}

	// создаем строку из массива слов с разделителем пробелом
	reverseS = strings.Join(words, " ")

	return reverseS
}
