package main

import "fmt"

type Vertex struct {
	X int 
	Y int
}
//隐式间接引用
func main(){
	v := Vertex{1,2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	//fmt.Println(p.X,(*p).x)
}