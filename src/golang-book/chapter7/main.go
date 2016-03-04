package main

import "fmt"

var counter int

func main() {
	// fmt.Println(counter)
	// oldCounter()
	// fmt.Println(counter)

	f := counterClosure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// fmt.Println(returnMultiple())
	// fmt.Println(returnNamed())
	// fmt.Println(variadricFunction(1, 2, 3))
	// closureFunction()
	// fmt.Println(halfInt(10))
	// fmt.Println(halfInt(9))
	// fmt.Println(variadricGreatestNumber(50, 10, 100, 1000, 50, 200))
	// f := fibonacci()
	// fmt.Println(f(), f(), f(), f(), f(), f(), f(), f(), f(), f())

}

func oldCounter() int {
	counter++
	return counter
}

func counterClosure() func() int {
	c := 0
	return func() int {
		c++
		return c
	}

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a = b
		b = a + b
		return a
	}
}

func returnNamed() (r int) {
	r = 2
	return
}

func returnMultiple() (x, y int) {
	x = 5
	y = 6
	return
}

func variadricFunction(args ...int) (total int) {
	total = 0
	for _, v := range args {
		total += v
	}
	return
}

func closureFunction() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}

func halfInt(x int) (result bool) {
	if (x % 2) != 0 {
		result = false
	} else {
		result = true
	}
	return
}

func variadricGreatestNumber(args ...int) (greatest int) {
	greatest = 0

	for _, v := range args {
		if greatest < v {
			greatest = v
		}
	}
	return
}
