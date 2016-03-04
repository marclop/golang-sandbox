package main

import "fmt"

func main() {
	// Basic var assignment
	var x string
	x = "Hello World"
	fmt.Println(x)

	// Variable overwrite
	x = "First"
	fmt.Println(x)
	x = "Second"
	fmt.Println(x)

	// Variable concatenation
	x = "First "
	fmt.Println(x)
	x = x + "Second"
	fmt.Println(x)

	// Variable (String) concatenation using special operator
	x = "First "
	fmt.Println(x)
	x += "Second"
	fmt.Println(x)

	// Comparisson result
	x = "hello"
	var y string
	y = "world"
	fmt.Println(x == y)
	y = "hello"
	fmt.Println(x == y)

	// Dynamic variable type
	w := 5
	fmt.Println(w)

	// Variable naming
	dogsName := "Max"
	fmt.Println("My dog's name is ", dogsName)
}
