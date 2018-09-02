package main

import "fmt"
//var p *int
func main(){
	i,j:= 41,42
	p :=&i
	fmt.Println(*p)
	*p =44
	fmt.Println(*p)

	p = &j
	*p = *p/2
	fmt.Println(j,*p)
}