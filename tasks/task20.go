package tasks

import (
	"regexp"
)

func Task20(s string) string {
	var reverseS string

	reWords, _ := regexp.Compile(`(\S)+`)
	words := reWords.FindAllString(s, -1)

	reSpaces, _ := regexp.Compile(`(\s)+`)
	spaces := reSpaces.FindAllString(s, -1)

	numberWords := len(words)

	for i := 0; i < numberWords; i++ {
		reverseS += words[numberWords-i-1]
		if i < numberWords-1 {
			reverseS += spaces[numberWords-i-2]
		}
	}

	return reverseS
}
