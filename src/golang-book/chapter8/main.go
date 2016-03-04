package main

import "fmt"

func main() {
	x := 5
	w := 0
	modifyPointer(&w)
	zeroNoPointer(x)
	fmt.Println(x) // x is 5
	zero(&x)
	fmt.Println(x) // x is 0
	// Access the memory position of var x
	fmt.Println(&x)
}
func zeroNoPointer(x int) {
	x = 0
}

// Modifies pointer
func zero(xPtr *int) {
	*xPtr = 0
}

func modifyPointer(wPtr *int) {
	w = 10
}
