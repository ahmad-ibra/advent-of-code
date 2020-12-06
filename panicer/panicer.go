package panicer

// Check function just panics :)
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
