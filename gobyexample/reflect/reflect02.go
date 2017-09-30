package main

import "fmt"
import "reflect"

func main() {
	var x float64 = 3.1
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("setability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("setability of v:", v.CanSet())
	v.SetFloat(4.1)
	fmt.Println(v.Interface())
	fmt.Println(v, x)
}
