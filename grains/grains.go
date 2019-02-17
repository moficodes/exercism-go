package grains

import "errors"

// Square calculates the number of grains
func Square(input int) (uint64, error) {
	if input > 64 || input < 1 {
		return 0, errors.New("Input invalid")
	}
	return 1 << (uint(input) - 1), nil
}

// Total calculate the total number of grains
func Total() uint64 {
	return (1 << 64) - 1
}
