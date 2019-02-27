package pascal

// Triangle returns the rows of a pascal trianlge
func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j < len(res[i]); j++ {
			if i == 0 {
				res[i][j] = 1
			} else if j == 0 {
				res[i][j] = res[i-1][j]
			} else if j == len(res[i])-1 {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}
