package math

// Package name
var Name string = "faca math package"

// Finds the average of a series of numbers
func Average(xs []float64) float64 {

	return Sum(xs) / float64(len(xs))
}
