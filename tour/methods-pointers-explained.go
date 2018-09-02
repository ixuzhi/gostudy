package main

import(
	"fmt"
	"math"
)
//指针与函数
type Vertex struct{
	X,Y float64
}

func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}

func Scale1(v Vertex,f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex,f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	v :=Vertex{3,4}
	Scale1(v,10)
	fmt.Println(v)
	Scale2(&v,10)
	fmt.Println(v)
}