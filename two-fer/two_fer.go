package twofer

import "strings"

const pattern = "One for you, one for me."

// ShareWith tells you how to share with the given name
func ShareWith(name string) string {
	result := pattern
	if name != "" {
		result = strings.Replace(result, "you", name, -1)
	}
	return result
}
