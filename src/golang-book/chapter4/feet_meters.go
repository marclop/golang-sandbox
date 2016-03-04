package main

import "fmt"

func main() {
	fmt.Println("This program will convert Feet to Meters")
	fmt.Print("Enter a feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input / 3.2808

	fmt.Println(output, "Meters")
}
