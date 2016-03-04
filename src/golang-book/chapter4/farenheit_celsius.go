package main

import "fmt"

func main() {
	fmt.Println("This program will convert Farenheit to Celisius")
	fmt.Print("Enter a farenheit temperature: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output, "ºC")
}
