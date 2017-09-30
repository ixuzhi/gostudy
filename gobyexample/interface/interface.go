package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width,heigh float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.heigh
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.heigh)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi *c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	fmt.Println("vim-go")
	r := rect{width:10,heigh:20}
	fmt.Println(r.area(),r.perim());
	
	c := circle{radius:10}

	fmt.Println(c.area(),c.perim())
	
	measure(r)
	fmt.Println()
	measure(c)

}
