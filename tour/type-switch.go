package main

import "fmt"
//i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type

func do(i interface{}){
	switch v := i.(type){
	case int:
		fmt.Printf("%v,%T\n",v,v)
	case string:
		fmt.Printf("%q,%T\n",v,v)
	default:
		fmt.Printf("%v,%T\n",v,v)
	}
}

func main(){
	do(1)
	do("hello")
	do(true)
}