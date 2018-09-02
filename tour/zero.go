package main
/*
没有明确初始值的变量声明会被赋予它们的 零值
*/
import "fmt"

func main(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i,f,b,s)
	fmt.Printf("%v %v %v %v\n",i,f,b,s)
}