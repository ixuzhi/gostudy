package main

import "fmt"
import "reflect"

func main() {
	fmt.Println("vim-go")
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64", v.Type(), v.Kind())
	fmt.Println(v.Type(), v.CanSet())
	v.SetFloat(4.1)
	fmt.Println(x)
}
