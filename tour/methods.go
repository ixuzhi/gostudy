package main

import(
	"fmt"
	"math"
)
//go 没有类，不过可以为结构体定义方法，方法就是一类带有特殊的接收者参数的函数
type Vertex struct {
	X,Y float64
}

func(v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
