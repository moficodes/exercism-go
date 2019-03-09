package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct takes digits and span and return a
// multiple of largest series or error if theres one
func LargestSeriesProduct(digits string, span int) (int64, error) {
	var res int64
	var mul int64 = 1
	if len(digits) < span || span < 0 {
		return -1, errors.New("invalid")
	}

	for i := 0; i <= len(digits)-span; i++ {
		for j := i; j < span+i; j++ {
			if !unicode.IsNumber(rune(digits[j])) {
				return -1, errors.New("wrong type")
			}
			mul *= int64(digits[j] - '0')
		}
		if mul > res {
			res = mul
		}
		mul = 1
	}
	return res, nil
}
