package main

import "fmt"

// Main function

func main() {
	// Int operations
	fmt.Println("1 + 1 =", 1+1)
	// Float operations
	fmt.Println("1 + 1 =", 1.0+1.0)
	// String operations
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	// Boolean operations
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(!true)
	// Multiply
	fmt.Println(321325 * 424521)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
