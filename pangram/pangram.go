package pangram

import (
	"unicode"
)

// IsPangram takes a string and return whether the string is pangram or not
func IsPangram(s string) bool {
	dict := make([]bool, 26)
	for _, v := range s {
		if unicode.IsLetter(v) {
			v = unicode.ToLower(v)
			dict[v-'a'] = true
		}
	}

	for _, v := range dict {
		if !v {
			return false
		}
	}
	return true
}
