// Package raindrops makes silly sounds if you feed it certain numbers.
package raindrops

import "strconv"

// TestVersion is so your tests and package don't get out of sync.
const (
	TestVersion = 1
	pling       = 1
	plang       = 2
	plong       = 4
)

// outcomes corresponding to bitmasked values
var outcomes = []string{
	"",           // 000
	"Pling",      // 001
	"Plang",      // 010
	"PlingPlang", // etc
	"Plong",
	"PlingPlong",
	"PlangPlong",
	"PlingPlangPlong",
}

// Convert turns a number into raindrop sounds if it has prime factors of 3, 5, or 7.
// Otherwise it returns the original number in string format, unchanged.
func Convert(n int) string {
	output := 0
	if n%5 == 0 {
		output = plang
	}
	if n%7 == 0 {
		output |= plong
	}
	if n%3 == 0 {
		return outcomes[output|pling]
	}
	if output == 0 {
		return strconv.Itoa(n)
	}
	return outcomes[output]
}
