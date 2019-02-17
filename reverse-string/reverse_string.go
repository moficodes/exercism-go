package reverse

// String Revers the string and return the result
func String(input string) string {
	a := []rune(input)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}
