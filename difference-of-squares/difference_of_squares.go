package diffsquares

// SquareOfSum Get the square of sum
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// SumOfSquares Get the sum of squares.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference Get the difference between the two
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
