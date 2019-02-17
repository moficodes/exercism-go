package hamming

import "errors"

// Distance finds the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length not equal")
	}
	count := 0
	ba := []byte(a)
	bb := []byte(b)

	for i := range ba {
		if ba[i] != bb[i] {
			count++
		}
	}
	return count, nil
}
