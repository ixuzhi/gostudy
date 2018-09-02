package main

import "fmt"
//类型通过实现 一个接口的所有方法来实现该接口
type I interface{
	M()
}

type T struct {
	S string
}
func( t T) M(){
	fmt.Println(t.S)
}

func main(){
	var i I = T{"hello"}
	i.M()
}