package main

import "fmt"

func main() {
	i, j := 42, 43
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
