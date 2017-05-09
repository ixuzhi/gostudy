package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"eric", "lily", "youta"}
	fmt.Println("vim-go")
	for _, name := range names {
		fmt.Printf("name_addr:%s,&name=%p,%#v\n", name, &name, name)
		go func() {
			fmt.Printf("hello,goroutine,name=%s,%p\n", name, &name)
		}()
	}
	fmt.Printf("name\n")
	runtime.Gosched()
}
