package erratum

// Use if a function of make believe land.
// Interesting though.
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	if res, err = o(); err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if f, ok := r.(FrobError); ok {
				res.Defrob(f.defrobTag)
			}
			err = r.(error)
		}
	}()

	res.Frob(input)
	return
}
