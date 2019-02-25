package accumulate

// Accumulate takes in an array and a function, returns the array
// after applying the function to every element
func Accumulate(data []string, converter func(string) string) []string {
	res := make([]string, len(data))
	for i, v := range data {
		res[i] = converter(v)
	}
	return res
}
