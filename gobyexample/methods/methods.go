package main

import "fmt"

type rect struct {
	width,heigh int
}

func (r *rect) area() int {
	return r.width * r.heigh
}

func (r rect) perim() int {
	return 2*(r.width + r.heigh)
}

func main() {
	fmt.Println("vim-go")
	r := rect{width:10,heigh:20}
	fmt.Println(r)
	fmt.Println("area:",r.area())

	fmt.Println("area:",r.perim())

	rp := &r

	fmt.Println("area:",rp.area())
	 
    fmt.Println("area:",rp.perim())
}
