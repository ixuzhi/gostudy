package main

import(
	"fmt"
	"math"
)
//函数值
//函数 也是值，可以作为函数的参数或者返回值
func compute(fn func(float64,float64)float64)float64{
	return fn(3,4)
}

func main(){
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}