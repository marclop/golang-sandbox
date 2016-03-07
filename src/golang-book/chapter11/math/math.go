package math

// Average Computes the average of an array
func Average(xs []float64) (total float64) {
	total = float64(0)
	for _, x := range xs {
		total += x
	}
	total = total / float64(len(xs))
	return
}

// Max Returns the maximum number inside a slice
func Max(xs []float64) (maximum float64) {
	maximum = 0
	for _, v := range xs {
		if maximum < v {
			maximum = v
		}
	}
	return
}

// Min Returns the maximum number inside a slice
func Min(xs []float64) (minimum float64) {
	minimum = 999999
	for _, v := range xs {
		if minimum > v {
			minimum = v
		}
	}
	return
}
