package main

import (
	"fmt"
	"golang-book/chapter11/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println("Average is", avg)
	fmt.Println("Minimum value is", min)
	fmt.Println("Maximum value is", max)
}
