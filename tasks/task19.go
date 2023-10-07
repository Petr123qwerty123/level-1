package tasks

func Task19(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		swapIndex := len(runes) - i - 1
		runes[i], runes[swapIndex] = runes[swapIndex], runes[i]
	}

	return string(runes)
}
