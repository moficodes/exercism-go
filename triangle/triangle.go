// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Exported for testing
const (
	// Pick values for the following identifiers used by the test program.
	NaT = 1 // not a triangle
	Equ = 2 // equilateral
	Iso = 3 // isosceles
	Sca = 4 // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = Sca
	if !isTraingle(a, b, c) {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || b == c || c == a {
		k = Iso
	}
	return k
}

func isTraingle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a >= math.Inf(1) || b >= math.Inf(1) || c >= math.Inf(1) {
		return false
	}
	return a+b >= c && b+c >= a && c+a >= b
}
