package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"eric", "lily", "youta"}
	fmt.Println("vim-go")
	for _, name := range names {
		go func(who string) {
			fmt.Printf("hello,goroutine,name=%s\n", who)
		}(name)
	}
	fmt.Printf("name\n")
	runtime.Gosched()
}
