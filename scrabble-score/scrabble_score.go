package scrabble

import "unicode"

var letterValues = [...]int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

// Score finds the scrabble score
func Score(word string) int {

	total := 0
	for _, v := range word {
		total += letterValues[unicode.ToLower(v)-'a']
	}
	return total
}
