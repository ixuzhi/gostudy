package main

import "fmt"
import "reflect"

func main() {
	type T struct {
		A int
		B string
	}
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d:%s %s=%v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
