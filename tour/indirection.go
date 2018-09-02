package main

import "fmt"
//带指针参数的函数必须接受一个指针
 //而以指针为接收者的方法被调用时，接收者既能为值又能为指针
 //Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)
 //接受一个值作为参数的函数必须接受一个指定类型的值
 //而以值为接收者的方法被调用时，接收者既能为值又能为指针
 /*
 	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // 编译错误！

	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // 编译错误！
 */
type Vertex struct{
	X,Y float64
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{3,4}
	v.Scale(2)
	fmt.Println(v)

	ScaleFunc(&v,10)
	fmt.Println(v)

	p :=&Vertex{3,4}
	p.Scale(2)
	fmt.Println(p)

	ScaleFunc(p,10)
	fmt.Println(p)



}