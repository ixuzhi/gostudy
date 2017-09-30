package main

import "fmt"

func f(from string) {
	for i:=0;i<3;i++ {
		fmt.Println(from,":",i)
	}
}
func main() {
	fmt.Println("vim-go")
	f("direct")

	go f("goroutiness")

	go func(msg string) {
		fmt.Println("msg:",msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done:",input)
}
