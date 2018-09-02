package main

import "fmt"
//make 动态创建
//make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
func main(){
	a := make([]int,5)
	printSlice("a",a)

	b :=make([]int,0,5)
	printSlice("b",b)

	c :=b[:2]
	printSlice("c",c)

	d :=c[2:5]
	printSlice("d",d)
}

func printSlice(s string,x []int){
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}