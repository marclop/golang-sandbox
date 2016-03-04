package main

import (
	"fmt"
	"math"
)

func main() {
	// short: c := Circle{0, 0, 5}
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	p := Person{"Marc"}
	p.Talk()
	robot := new(Robot)
	robot.Name = "Bender"
	robot.Talk()
}

// Circle shape
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle shape
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// Person Struct
type Person struct {
	Name string
}

// Talk outputs to stdout
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Robot struct
type Robot struct {
	Person
	Model string
}
