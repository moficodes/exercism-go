package etl

// Transform takes in a map of int to string array and return a
func Transform(input map[int][]string) map[string]int {
	res := make(map[string]int)
	for k, v := range input {
		for _, d := range v {
			res[d] = k
		}
	}
	return res
}
