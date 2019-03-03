package summultiples

// SumMultiples adds up all the uniqe multiples of the divisors
func SumMultiples(limit int, divisors ...int) int {
	res := 0
	multiples := make(map[int]struct{})

	if len(divisors) == 0 {
		return res
	}
	count := 1
	for _, v := range divisors {
		for v != 0 && v*count < limit {
			multiples[v*count] = struct{}{}
			count++
		}
		count = 1
	}

	for k := range multiples {
		res += k
	}
	return res
}
