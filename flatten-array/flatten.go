package flatten

func Flatten(in interface{}) []interface{} {
	output := []interface{}{}
	switch e := in.(type) {
	case []interface{}:
		for _, v := range e {
			output = append(output, Flatten(v)...)
		}
	case interface{}:
		output = append(output, e)
	}
	return output
}
