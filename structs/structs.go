package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(person{"bob",20})
	fmt.Println(person{name:"bobo2",age:18})
	fmt.Println(person{name:"bob3"})
	fmt.Println(&person{name:"bob4",age:19})
	s := person{name:"bob5",age:21}
	fmt.Println(s)
	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)
	fmt.Println(s.age)
	sp.age = 18
	sp.name = "xuxu"
	
	fmt.Println(s)

	fmt.Println(sp)
}
