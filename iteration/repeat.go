package iteration

// Repeat some string N times
func Repeat(character string, repetitions int) string {
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += character
	}
	return repeated
}
