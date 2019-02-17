package isogram

import "unicode"

// IsIsogram Check if a given string is Isogram
func IsIsogram(input string) bool {
	table := make([]bool, 26)
	for _, r := range input {
		if unicode.IsLetter(r) {
			val := unicode.ToLower(r) - 'a'
			if table[val] {
				return false
			}
			table[val] = true
		}
	}
	return true
}
