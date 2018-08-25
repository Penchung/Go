package math

// Finds the sum of a series of numbers
func Sum(xs []float64) (total float64) {

	for _, x := range xs {
		total += x
	}
	return
}
