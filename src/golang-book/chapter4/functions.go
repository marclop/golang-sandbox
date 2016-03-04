package main

import "fmt"

var x = "Hello World"

const y = "Constant value"

var (
	a = 5
	b = 6
	c = 7
)

func main() {
	printVariable()
	printConstant()
}

func printVariable() {
	fmt.Println(x)
}

func printConstant() {
	fmt.Println(y)
}
