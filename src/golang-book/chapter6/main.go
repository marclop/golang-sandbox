package main

import "fmt"

/*
  Arrays, Slices and Maps
*/
func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	arrays()
	slices()
	maps()
	nestedMaps()
	fmt.Println(getSmallestNumber(x))
}

func arrays() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	//var total float64
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func slices() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice2, slice3, slice2[3])
}

func maps() {
	// Create maps with make (no value required)
	x := make(map[string]int)
	x["key"] = 10
	y := make(map[int]int)
	y[1] = 50
	fmt.Println(x["key"])
	fmt.Println(y[1])
	// Delete map value
	delete(y, 1)
	delete(x, "key")

	elements := map[string]string{
		"H":  "Hidrogen",
		"He": "Helium",
	}

	var elementList = [2]string{
		"He",
		"Un",
	}
	// If the element exists return the name of the element and true
	for _, value := range elementList {
		if name, ok := elements[value]; ok {
			fmt.Println(name, ok)
		} else {
			fmt.Println("Element not found")
		}
	}

}

func nestedMaps() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

// Named return
func getSmallestNumber(xs []int) (smallest int) {
	// Named returns don't work with :=
	smallest = xs[0]

	for _, v := range xs {
		if smallest > v {
			smallest = v
		}
	}

	return
}
