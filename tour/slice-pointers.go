package main

import "fmt"
//a[low : high]它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
//切片不存储任何数据，修改切片会修改实际数据
func main(){
	names :=[4]string{
		"johb",
		"frank",
		"alice",
		"jack",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:2]
	fmt.Println(a,b)

	b[0] = "xxx"
	fmt.Println(a,b)
	fmt.Println(names)
}