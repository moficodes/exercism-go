package secret

// Handshake takes an unsigned integer and returns a string array
func Handshake(n uint) []string {
	res := make([]string, 0, 5)
	if n&1 != 0 {
		res = append(res, "wink")
	}
	if n&2 != 0 {
		res = append(res, "double blink")
	}
	if n&4 != 0 {
		res = append(res, "close your eyes")
	}
	if n&8 != 0 {
		res = append(res, "jump")
	}
	if n&16 != 0 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}
