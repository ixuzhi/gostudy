package main

import (
	"fmt"
	"runtime"
)

func main() {
	name := "eric"
	fmt.Println("vim-go")
	go func() {
		fmt.Printf("hello,goroutine,name=%s\n", name)
	}()
	name = "harry"
	fmt.Printf("name:%s\n", name)
	runtime.Gosched()
}
