package main

import "fmt"

func main(){
	defer fmt.Println("defer world")
	fmt.Println("hello")
}